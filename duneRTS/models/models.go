package models

// SpiceHarvester type - a critical first build for taming the desert planet
type SpiceHarvester struct {
    Speed int
    Full bool
    WormSign bool
    Crew int
}


// if Speed of SpiceHarvester is 5 or below
// return true
func (h SpiceHarvester) IsWorried() (bool){
    return h.Speed <= 5
}


func (h *SpiceHarvester) AddCrew() {
    h.Crew++
}
