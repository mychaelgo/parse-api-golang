package parseAPI

type Push struct {
	Where					Where			`json:"where"`
	Data					Data			`json:"data"`
}

type Where struct {
	InstalationId			string			`json:"installationId"`
}

type Data struct {
	Alert					string			`json:"alert"`
}