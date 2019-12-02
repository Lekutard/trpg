package ports

//GeneralRollInputPort 一般判定InputPort
type GeneralRollInputPort struct {
	TargetValue int
	Ajust       int
}

//GeneralRollOutputPort 一般判定InputPort
type GeneralRollOutputPort struct {
	Success    bool
	Critical   bool
	Fumble     bool
	Difference int
}
