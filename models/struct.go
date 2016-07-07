package models

//Phystopo - A struct to hold subotops
type Phystopo struct {
	Name     string
	Kilo     []string
	Liberty  []string
	Rdo      []string
	El7      []string
	Ubuntu14 []string
}

//RunDB - A struct to hold Run info
type RunDB struct {
	ID            int    `json:"id,string"`
	TimeStart     string `json:"time_start"`
	LoadID        string `json:"load_id"`
	TestbedID     int    `json:"testbed_id,string"`
	RequestedBy   string `json:"requested_by"`
	Link          string `json:"link"`
	RunLevel      int    `json:"run_level,string"`
	RunType       int    `json:"run_type,string"`
	TcPassed      int    `json:"tc_passed,string"`
	TcFailed      int    `json:"tc_failed,string"`
	TcSkipped     int    `json:"tc_skipped,string"`
	TcPp          int    `json:"tc_pp,string"`
	TcPassedPrev  int    `json:"tc_passed_prev,string"`
	TimeDuration  int    `json:"time_duration,string"`
	TimeEstimate  int    `json:"time_estimate,string"`
	TimeSubmit    string `json:"time_submit"`
	Jobid         string `json:"jobid"`
	Jobstate      int    `json:"jobstate,string"`
	Args          string `json:"args"`
	Flag          int    `json:"flag,string"`
	LoadhistoryID int    `json:"loadhistory_id,string"`
	CurrentTestID int    `json:"current_test_id,string"`
}

//Build - ID and Load name
type Build struct {
	ID   int    `json:"id,string"`
	Load string `json:"load"`
}
