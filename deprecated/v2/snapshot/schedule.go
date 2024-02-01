package snapshot

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Schedule stores the time, destination, and number of snapshots to keep. All
// the time values are represented by pointers, and set only when used.
type Schedule struct {
	Destination Destination   `json:"destination"`
	Keep        int           `json:"keep"`
	Minutely    *float32      `json:"minutely,omitempty"`
	Minute      *int          `json:"minute,omitempty"`
	Hour        *int          `json:"hour,omitempty"`
	Weekday     *time.Weekday `json:"weekday,omitempty"`
	Day         *int          `json:"day,omitempty"`
	Month       *time.Month   `json:"month,omitempty"`
	YearDay     *int          `json:"yearday,omitempty"`
	Options     Option        `json:"options,omitempty"`
}

//go:generate go run "github.com/dmarkham/enumer" -type Option -transform title-upper -text

// Option is a bitmask of settings for a Schedule.
type Option int

// A Schedule may have one or more of the following Options.
const (
	PauseVM Option = 1 << iota
)

// Schedules is a set of snapshot schedules over every possible timing Period.
type Schedules []Schedule

// NewSchedule returns and initialized Schedule. Use the option setting
// functions to set the time and other fields.
func NewSchedule(options ...func(*Schedule)) Schedule {
	s := Schedule{Keep: 1}

	for _, opts := range options {
		opts(&s)
	}

	s.normalize()

	return s
}

// EveryMinute is an option setting function for the Schedule constructor. This
// sets the schedule to match on every minute.
func EveryMinute() func(*Schedule) {
	return func(s *Schedule) {
		val := float32(1)
		s.Minutely = &val
	}
}

// Minute is an option setting function for the Schedule constructor. It sets
// the minute used by an Hourly snapshot.
func Minute(minute int) func(*Schedule) {
	return func(s *Schedule) {
		s.Minute = &minute
	}
}

// Hour is an option setting function for the Schedule constructor. It sets the
// hour used by a Daily snapshot.
func Hour(hour int) func(*Schedule) {
	return func(s *Schedule) {
		s.Hour = &hour
	}
}

// Weekday is an option setting function for the Schedule constructor. It sets
// the weekday used by a Weekly snapshot.
func Weekday(weekday time.Weekday) func(*Schedule) {
	return func(s *Schedule) {
		s.Weekday = &weekday
	}
}

// Day is an option setting function for the Schedule constructor. It sets the
// day used by a Monthly snapshot.
func Day(day int) func(*Schedule) {
	return func(s *Schedule) {
		s.Day = &day
	}
}

// Month is an option setting function for the Schedule constructor. It sets the
// month used by a Yearly snapshot.
func Month(month time.Month) func(*Schedule) {
	return func(s *Schedule) {
		s.Month = &month // mutually exclusive
		s.YearDay = nil
	}
}

// YearDay is an option setting function for the Schedule constructor. It sets
// the day of the year used by a Yearly snapshot.
func YearDay(day int) func(*Schedule) {
	return func(s *Schedule) {
		s.Month = nil // mutually exclusive
		s.YearDay = &day
	}
}

// Keep is an option setting function for the Schedule constructor. It sets the
// number of snapshots to keep.
func Keep(n int) func(*Schedule) {
	return func(s *Schedule) {
		s.Keep = n
	}
}

// IsTime returns true if the Schedule matches t.
func IsTime(s Schedule, t time.Time) bool {
	return s.IsTime(t)
}

// Matches returns true if both schedule are for the same time. It does not
// compare the Destination or Keep values.
func (s Schedule) Matches(t *Schedule) bool {
	p := s.Period()
	if p != t.Period() {
		return false
	}

	switch p {
	case Minutely:
		return true
	case Hourly:
		return s.matchesMinute(t)
	case Daily:
		return s.matchesMinute(t) && s.matchesHour(t)
	case Weekly:
		return s.matchesMinute(t) && s.matchesHour(t) && s.matchesWeekday(t)
	case Monthly:
		return s.matchesMinute(t) && s.matchesHour(t) && s.matchesDay(t)
	case Yearly:
		return s.matchesMinute(t) && s.matchesHour(t) && (s.matchesMonth(t) || s.matchesYearDay(t))
	}

	return false
}

// IsTime returns true if the SnapShotTime matches t.
func (s Schedule) IsTime(t time.Time) bool {
	isMinute := s.isMinute(t)
	isHour := s.isHour(t)
	isWeekday := s.isWeekday(t)
	isDay := s.isDay(t)
	isMonth := s.isMonth(t)
	isYearDay := s.isYearDay(t)

	switch s.Period() {
	case Yearly: // can be a day of year or a month
		if isYearDay {
			return isHour && isMinute
		}
		return isMonth && isDay && isHour && isMinute
	case Monthly:
		return isDay && isHour && isMinute
	case Weekly:
		return isWeekday && isHour && isMinute
	case Daily:
		return isHour && isMinute
	case Hourly:
		return isMinute
	case Minutely: // [0,1) probability of triggering
		if s.Minutely != nil && rand.Float32() <= *s.Minutely { // nolint
			return true
		}
	}

	return false
}

// Period returns the SnapshotPeriod of s. This is determined from which set of
// times are set.
func (s Schedule) Period() Period {
	switch { // order is critical here don't reshuffle cases
	case s.YearDay != nil || s.Month != nil:
		return Yearly
	case s.Day != nil:
		return Monthly
	case s.Weekday != nil:
		return Weekly
	case s.Hour != nil:
		return Daily
	case s.Minute != nil:
		return Hourly
	case s.Minutely != nil:
		return Minutely
	}

	return -1
}

// when is defined as a type only to provide a custom String() method.
type when int

// String implements the stringer interface for w. A invalid value of -1 returns
// the empty string.
func (w when) String() string {
	if w == -1 {
		return ""
	}

	return strconv.Itoa(int(w))
}

// When returns the time of the Schedule as a number in the Period. For example:
// an Hourly schedule returns the minute of the hour.
func (s Schedule) When() when { // nolint:revive
	switch {
	case s.YearDay != nil:
		return when(*s.YearDay)
	case s.Month != nil:
		return when(*s.Month)
	case s.Day != nil:
		return when(*s.Day)
	case s.Weekday != nil:
		return when(*s.Weekday)
	case s.Hour != nil:
		return when(*s.Hour)
	case s.Minute != nil:
		return when(*s.Minute)
	case s.Minutely != nil:
		return when(*s.Minutely * 100)
	}

	return -1
}

// JSON returns the json representation of the Schedule.
func (s Schedule) JSON() []byte {
	bytes, err := json.Marshal(s)
	if err != nil {
		return []byte("<nil>")
	}

	return bytes
}

// String implements the Stringer interface.
func (s Schedule) String() string {
	str := &strings.Builder{}

	switch s.Period() {
	case Minutely:
		if s.Minutely != nil && *s.Minutely > 0 {
			freq := int(*s.Minutely * 100) // optional probability [0-100)
			fmt.Fprintf(str, "%v=%v@%d", s.Destination, s.Keep, freq)
			break
		}

		fmt.Fprintf(str, "%v=%v", s.Destination, s.Keep)
	case Hourly:
		if s.Minute != nil && *s.Minute > 0 { // normally there is no Minute just run at top of hour
			fmt.Fprintf(str, "%v=%v@%v", s.Destination, s.Keep, *s.Minute)
			break
		}

		fmt.Fprintf(str, "%v=%v", s.Destination, s.Keep)
	case Daily:
		var hour int

		if s.Hour != nil {
			hour = *s.Hour
		}

		fmt.Fprintf(str, "%v=%v@%v", s.Destination, s.Keep, hour)
	case Weekly:
		var day time.Weekday

		if s.Weekday != nil {
			day = *s.Weekday
		}

		fmt.Fprintf(str, "%v=%v@%v", s.Destination, s.Keep, day.String()[0:3])
	case Monthly:
		var day int

		if s.Day != nil {
			day = *s.Day
		}

		fmt.Fprintf(str, "%v=%v@%v", s.Destination, s.Keep, day)
	case Yearly:
		var when string

		if s.Month != nil {
			when = s.Month.String()[0:3]
		}

		if s.YearDay != nil && *s.YearDay > 0 {
			when = fmt.Sprintf("%03d", *s.YearDay)
		}

		fmt.Fprintf(str, "%v=%d@%v", s.Destination, s.Keep, when)
	}

	fmt.Fprint(str, s.Options.strings())

	return str.String()
}

func (o Option) strings() string {
	if o == 0 {
		return ""
	}

	// At one time the options were a list, so keep the internals like that.
	// Right now PauseVM is the only option.

	var s []string

	if o&PauseVM > 0 {
		s = append(s, PauseVM.String())
	}

	return fmt.Sprintf(" options:%s", strings.Join(s, ","))
}

// UnmarshalJSON implements the Unmarshaler interface for s. Any missing time
// periods will get their default values. For example, a Hour snapshot with a
// missing Minute field will have Minute set to the top of the hour (0).
func (s *Schedule) UnmarshalJSON(text []byte) error {
	type alias Schedule
	var a alias

	err := json.Unmarshal(text, &a)
	if err != nil {
		return err
	}

	*s = Schedule(a)
	s.normalize()

	return nil
}

// Parse parses the chronish lines for scheduling snapshot. Parsing is line
// oriented and the text can use "#" as comments. An example of the line format:
//
//	hourly:local=10,archive=3 daily:local=3,remote=3@10,remote=1@12 weekly:local=1@Tue [options:...]
//
// See sched_test.go for more examples of the format.
func Parse(text string) (Schedules, error) {
	times := make([]Schedule, 0)

	for _, line := range strings.Split(text, "\n") {
		i := strings.Index(line, "#") // strip out comments
		if i >= 0 {
			line = line[:i]
		}

		var options Option

		for _, schedule := range strings.Fields(line) { // split into hourly,daily,... schedules
			i := strings.Index(schedule, ":")
			if i <= 0 {
				continue // not a schedule
			}

			if schedule[:i] == "options" {
				if len(schedule) > i+1 {
					var err error
					options, err = parseOptions(schedule[i+1:])
					if err != nil {
						return nil, err
					}
				}
				continue // empty options
			}

			t, err := parse(schedule)
			if err != nil {
				return nil, err
			}

			times = append(times, t...)
		}

		// Options bind to every snapshot time parsed on the current line.
		if options > 0 {
			for i := range times {
				times[i].Options = options
			}

		}
	}

	return times, nil
}

func parseOptions(text string) (Option, error) {
	var options Option

	for _, o := range strings.Split(text, ",") {
		o, err := OptionString(o)
		if err != nil {
			return 0, err
		}
		options |= o
	}

	return options, nil
}

// parse parses a snapshot string of the form:
//
//	period:destination=keep[@when][,destination=keep[@when]...]
//
// for example:
//
//	hourly:local=3@30
//
// Only a single period (e.g "hourly") can be present, but multiple snapshots
// within that period are allowed.
func parse(schedule string) ([]Schedule, error) {
	i := strings.Index(schedule, ":")
	if i <= 0 {
		return nil, fmt.Errorf("not a schedule: %q", schedule)
	}

	period, err := PeriodString(schedule[:i])
	if err != nil {
		return nil, fmt.Errorf("line %q: %w", schedule, err)
	}

	// NOTE: Disable Minutely snapshots for Production. Build tags protect some of
	// this check is needed as well. Silently ignore these Schedules.

	if period == Minutely && isProduction {
		return nil, nil
	}

	if i == len(schedule)+1 { // only "period:" with no data
		return nil, fmt.Errorf("invalid schedule %q", schedule)
	}

	// Loop over all the snapshots for the given period and parse them.
	// The snapshot string is of the form destination=keep[@when].

	snapshots := strings.Split(schedule[i+1:], ",") // i+1 to skip the ":"
	times := make([]Schedule, 0, len(snapshots))

	for _, snapshot := range snapshots {
		s, err := parseSchedule(period, snapshot)
		if err != nil {
			return nil, err
		}

		times = append(times, *s)
	}

	return times, nil
}

// Minutely returns a slice of Minutely snapshots in the schedule.
func (s Schedules) Minutely() Schedules {
	return s.match(Minutely)
}

// Hourly returns a slice of Hourly snapshots in the schedule.
func (s Schedules) Hourly() Schedules {
	return s.match(Hourly)
}

// Daily returns a slice of Daily snapshots in the schedule.
func (s Schedules) Daily() Schedules {
	return s.match(Daily)
}

// Weekly returns a slice of Weekly snapshots in the schedule.
func (s Schedules) Weekly() Schedules {
	return s.match(Weekly)
}

// Monthly returns a slice of Monthly snapshots in the schedule.
func (s Schedules) Monthly() Schedules {
	return s.match(Monthly)
}

// Yearly returns a slice of Yearly snapshots in the schedule.
func (s *Schedules) Yearly() Schedules {
	return s.match(Yearly)
}

// Sort sorts the Schedules. This is used by the String() method.
func (s Schedules) Sort() {
	sort.SliceStable(s, func(i, j int) bool {
		switch {
		case s[i].YearDay != nil && s[j].YearDay != nil:
			return *s[i].YearDay < *s[j].YearDay
		case s[i].Month != nil && s[j].Month != nil:
			return *s[i].Month < *s[j].Month
		case s[i].Day != nil && s[j].Day != nil:
			return *s[i].Day < *s[j].Day
		case s[i].Weekday != nil && s[j].Weekday != nil:
			return *s[i].Weekday < *s[j].Weekday
		case s[i].Hour != nil && s[j].Hour != nil:
			return *s[i].Hour < *s[j].Hour
		case s[i].Minute != nil && s[j].Minute != nil:
			return *s[i].Minute < *s[j].Minute
		}

		return false
	})
}

// String implements the fmt.Stringer interface for the SnapshotSchedule.
func (s Schedules) String() string {
	var schedule, minutely, hourly, daily, weekly, monthly, yearly []string

	s.Sort()

	for _, snap := range s {
		switch snap.Period() {
		case Minutely:
			minutely = append(minutely, snap.String())
		case Hourly:
			hourly = append(hourly, snap.String())
		case Daily:
			daily = append(daily, snap.String())
		case Weekly:
			weekly = append(weekly, snap.String())
		case Monthly:
			monthly = append(monthly, snap.String())
		case Yearly:
			yearly = append(yearly, snap.String())
		}
	}

	if minutely != nil {
		schedule = append(schedule, Minutely.String()+":"+strings.Join(minutely, ","))
	}

	if hourly != nil {
		schedule = append(schedule, Hourly.String()+":"+strings.Join(hourly, ","))
	}

	if daily != nil {
		schedule = append(schedule, Daily.String()+":"+strings.Join(daily, ","))
	}

	if weekly != nil {
		schedule = append(schedule, Weekly.String()+":"+strings.Join(weekly, ","))
	}

	if monthly != nil {
		schedule = append(schedule, Monthly.String()+":"+strings.Join(monthly, ","))
	}

	if yearly != nil {
		schedule = append(schedule, Yearly.String()+":"+strings.Join(yearly, ","))
	}

	return strings.Join(schedule, " ")
}

// normalize sets the default values for any missing time periods. For example a
// Monthly snapshot with nil values for the hour, minute, or day will set those
// to 0, 0, and 1 respectively.
//
// A Schedule MUST call this prior to any use, all operations other than
// parsing assume the Schedule is complete down to the minute of the hour it
// is to run.
func (s *Schedule) normalize() {
	zero := 0
	one := 1

	if s.YearDay != nil || s.Month != nil || s.Day != nil || s.Weekday != nil {
		if s.Hour == nil {
			s.Hour = &zero
		}

		if s.Minute == nil {
			s.Minute = &zero
		}
	}

	if s.Hour != nil {
		if s.Minute == nil {
			s.Minute = &zero
		}
	}

	if s.Month != nil {
		if s.Day == nil {
			s.Day = &one
		}
	}
}

func (s Schedule) matchesMinute(t *Schedule) bool {
	switch {
	case s.Minute == nil && t.Minute == nil:
		return true
	case s.Minute != nil && t.Minute != nil && *s.Minute == *t.Minute:
		return true
	default:
		return false
	}
}

func (s Schedule) matchesHour(t *Schedule) bool {
	switch {
	case s.Hour == nil && t.Hour == nil:
		return true
	case s.Hour != nil && t.Hour != nil && *s.Hour == *t.Hour:
		return true
	default:
		return false
	}
}

func (s Schedule) matchesWeekday(t *Schedule) bool {
	switch {
	case s.Weekday == nil && t.Weekday == nil:
		return true
	case s.Weekday != nil && t.Weekday != nil && *s.Weekday == *t.Weekday:
		return true
	default:
		return false
	}
}

func (s Schedule) matchesDay(t *Schedule) bool {
	switch {
	case s.Day == nil && t.Day == nil:
		return true
	case s.Day != nil && t.Day != nil && *s.Day == *t.Day:
		return true
	default:
		return false
	}
}

func (s Schedule) matchesMonth(t *Schedule) bool {
	switch {
	case s.Month == nil && t.Month == nil:
		return true
	case s.Month != nil && t.Month != nil && *s.Month == *t.Month:
		return true
	default:
		return false
	}
}

func (s Schedule) matchesYearDay(t *Schedule) bool {
	switch {
	case s.YearDay == nil && t.YearDay == nil:
		return true
	case s.YearDay != nil && t.YearDay != nil && *s.YearDay == *t.YearDay:
		return true
	default:
		return false
	}
}

func (s Schedule) isMinute(t time.Time) bool {
	return s.Minute != nil && *s.Minute == t.Minute()
}

func (s Schedule) isHour(t time.Time) bool {
	return s.Hour != nil && *s.Hour == t.Hour()
}

func (s Schedule) isWeekday(t time.Time) bool {
	return s.Weekday != nil && *s.Weekday == t.Weekday()
}

func (s Schedule) isDay(t time.Time) bool {
	return s.Day != nil && *s.Day == t.Day()
}

func (s Schedule) isMonth(t time.Time) bool {
	return s.Month != nil && *s.Month == t.Month()
}

func (s Schedule) isYearDay(t time.Time) bool {
	return s.YearDay != nil && *s.YearDay == t.YearDay()
}

// parseSchedule parses a single Schedule. The period is already
// stripped off the string. The string is of the form:
//
//	dest=keep[@when]
func parseSchedule(period Period, text string) (*Schedule, error) {
	i := strings.Index(text, "=")
	if i <= 0 {
		return nil, fmt.Errorf("keep count not specified %q", text)
	}

	dest, err := DestinationString(text[:i])
	if err != nil {
		return nil, fmt.Errorf("snapshot %q:%w", text, err)
	}

	if i == len(text)+1 { // only "destination=" with no data
		return nil, fmt.Errorf("invalid snapshot %q", text)
	}

	var (
		keep int
		when string
	)

	t := strings.Split(text[i+1:], "@") // keep@when or keep
	if len(t) > 2 {
		return nil, fmt.Errorf("invalid snapshot %q", text)
	}

	keep, err = strconv.Atoi(t[0])
	if err != nil {
		return nil, fmt.Errorf("snapshot %q: %w", text, err)
	}

	if len(t) > 1 {
		when = t[1]
	}

	s := Schedule{
		Destination: dest,
		Keep:        keep,
	}

	if err := s.setWhen(period, when); err != nil {
		return nil, err
	}

	s.normalize()

	return &s, nil
}

// setWhen adds a schedule time based on the period and the when string parsing
// rules.
func (s *Schedule) setWhen(period Period, when string) error {
	var err error

	switch period {
	case Minutely:
		err = s.setMinutely(when)
	case Hourly:
		err = s.setHourly(when)
	case Daily:
		err = s.setDaily(when)
	case Weekly:
		err = s.setWeekly(when)
	case Monthly:
		err = s.setMonthly(when)
	case Yearly:
		err = s.setYearly(when)
	}

	if err != nil {
		return err
	}

	return nil
}

// setMinutely sets the schedule to trigger every minute. The when parameter
// should be a number from 0-100 indicated the percentage chance of the schedule
// firing on an given minute. For example, a when of 20 means every 5 minutes.
func (s *Schedule) setMinutely(when string) error {
	if when == "" {
		when = "100"
	}

	n, err := strconv.Atoi(when)
	if err != nil {
		return err
	}

	switch {
	case n < 0:
		n = 0
	case n > 100:
		n = 100
	}

	x := float32(n) / 100
	s.Minutely = &x

	return nil
}

// setHourly sets the Minute of the snapshot based on the when string.
//
// Hourly Time Options:
//
//	Minute of the Hour: 0-59
//	Minute of the Hour, padded with zeros: 00-59
func (s *Schedule) setHourly(when string) error {
	if when == "" {
		when = "0"
	}

	if len(when) > 1 && when[0:1] == "0" {
		when = when[1:]
	}

	minute, err := strconv.Atoi(when)
	if err != nil {
		return err
	}

	if !(minute >= 0 && minute <= 59) {
		return fmt.Errorf("invalid minute %d", minute)
	}

	s.Minute = &minute

	return nil
}

// setDaily sets the hour of the snapshot based on the when string.
//
// Daily Time Options:
//
//	Hour in 24-hour time: 0-23
//	Hour in 24-hour time, padded with zeros: 00-23
func (s *Schedule) setDaily(when string) error {
	if when == "" {
		when = "0"
	}

	if len(when) > 1 && when[0:1] == "0" {
		when = when[1:]
	}

	hour, err := strconv.Atoi(when)
	if err != nil {
		return err
	}

	if !(hour >= 0 && hour <= 23) {
		return fmt.Errorf("invalid hour %d", hour)
	}

	s.Hour = &hour

	return nil
}

// setWeeks sets the Weekday of the snapshot based on the when string.
//
// Weekly Time Options:
//
//	Number Day of the Week: 0-6 (Starting with Sunday)
//	Abbreviated Name of the Day of the Week: Sun,Mon,Tue,Wed,Thu,Fri,Sat
//	Full Name of Day of the Week: Sunday,Monday,Tuesday,etc.
func (s *Schedule) setWeekly(when string) error {
	var day time.Weekday

	if when == "" {
		when = time.Sunday.String()
	}

	n, err := strconv.Atoi(when)
	if err == nil {
		day = time.Weekday(n)
	} else {
		var ok bool

		if len(when) > 3 {
			when = when[:3]
		}

		day, ok = weekdays[strings.ToLower(when)]
		if !ok {
			return fmt.Errorf("invalid weekday %q %v", when, weekdays)
		}
	}

	if !(day >= time.Sunday && day <= time.Saturday) {
		return fmt.Errorf("invalid day %v", day)
	}

	s.Weekday = &day

	return nil
}

// setMonthly sets the Day of the snapshot based on the when string.
//
// Monthly Time Options:
//
//	Number Day of the Month: 1-31
//	Number Day of the Month, padded with zeros: 01-31
func (s *Schedule) setMonthly(when string) error {
	if when == "" {
		when = "1"
	}

	if len(when) > 1 && when[0:1] == "0" {
		when = when[1:]
	}

	day, err := strconv.Atoi(when)
	if err != nil {
		return err
	}

	if !(day >= 1 && day <= 31) {
		return fmt.Errorf("invalid day %d", day)
	}

	s.Day = &day

	return nil
}

// setYearly sets either the Month or the YearDay of the snapshot based on the when string.
//
// Yearly Time Options:
//
//	Number Month of the Year, padded with zeros: 01-12
//	Day of the Year, padded with zeros: 001-366
//	Abbreviated Name of Month: Jan,Feb,Mar,Apr,May,Jun,Jul,Aug,Sep,Oct,Nov,Dec
//	Full Name of the Month: January,February,March,April,etc.
func (s *Schedule) setYearly(when string) error {
	var (
		month   time.Month
		yearday int
		ok      bool
	)

	if when == "" {
		when = time.January.String()
	}

	if len(when) >= 3 {
		month, ok = months[strings.ToLower(when[:3])]
	}

	if !ok {
		length := len(when)

		if length >= 2 && when[0:1] == "0" { // remove padded zeros
			when = when[1:]
		}
		if len(when) >= 2 && when[0:1] == "0" {
			when = when[1:]
		}

		n, err := strconv.Atoi(when)
		if err != nil {
			return fmt.Errorf("invalid day %q", when)
		}

		switch length {
		case 2:
			month = time.Month(n)
		case 3:
			yearday = n
		default:
			return fmt.Errorf("invalid day %q", when)
		}
	}

	if month != 0 {
		if !(month >= time.January && month <= time.December) {
			return fmt.Errorf("invalid month %v", month)
		}

		day := 1
		s.Day = &day
		s.Month = &month
	}

	if yearday != 0 {
		if !(yearday >= 1 && yearday <= 366) {
			return fmt.Errorf("invalid yearday %d", yearday)
		}

		s.YearDay = &yearday
	}

	return nil
}

func (s Schedules) match(period Period) Schedules {
	var times Schedules

	for _, snap := range s {
		if snap.Period() == period {
			times = append(times, snap)
		}
	}

	return times
}

var (
	weekdays map[string]time.Weekday
	months   map[string]time.Month
)

func init() {
	weekdays = make(map[string]time.Weekday)
	for i := 0; i <= 7; i++ {
		d := time.Weekday(i)
		s := strings.ToLower(d.String()[:3])
		weekdays[s] = d
	}

	months = make(map[string]time.Month)
	for i := 0; i <= 12; i++ {
		m := time.Month(i)
		s := strings.ToLower(m.String()[:3])
		months[s] = m
	}

}
