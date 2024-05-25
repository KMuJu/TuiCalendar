package tui

func NewWeekTable() WeekTable {
	week := WeekTable{}
	return week
}

func (week *WeekTable) Render() string {
	return week.Table.Render()
}

func (week *WeekTable) Update() {

}
