package main

type Peoples interface {
	Get() float64
}

type People []struct {
	Active  bool `json:"active"`
	Address []struct {
		City     string `json:"city"`
		Country  string `json:"country"`
		District string `json:"district"`
		ID       string `json:"id"`
		Line     string `json:"line"`
		Period   struct {
		} `json:"period"`
		PostalCode string `json:"postalCode"`
		State      string `json:"state"`
		Text       string `json:"text"`
		Type       string `json:"type"`
		Use        string `json:"use"`
	} `json:"address"`
	BirthDate     string `json:"birthDate"`
	Communication []struct {
		ID       string `json:"id"`
		Language struct {
		} `json:"language"`
		Preferred bool `json:"preferred"`
	} `json:"communication"`
	Contact []struct {
		Address struct {
		} `json:"address"`
		Gender string `json:"gender"`
		ID     string `json:"id"`
		Name   struct {
		} `json:"name"`
		Organization struct {
		} `json:"organization"`
		Period struct {
		} `json:"period"`
		Relationship []struct {
			Coding []struct {
				Code         string `json:"code"`
				Display      string `json:"display"`
				ID           string `json:"id"`
				System       string `json:"system"`
				UserSelected bool   `json:"userSelected"`
				Version      string `json:"version"`
			} `json:"coding"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"relationship"`
		Telecom []struct {
			ID     string `json:"id"`
			Period struct {
			} `json:"period"`
			Rank   int    `json:"rank"`
			System string `json:"system"`
			Use    string `json:"use"`
			Value  string `json:"value"`
		} `json:"telecom"`
	} `json:"contact"`
	Contained []struct {
		ID            string `json:"id"`
		ImplicitRules string `json:"implicitRules"`
		Language      string `json:"language"`
		Meta          struct {
		} `json:"meta"`
	} `json:"contained"`
	Deceased            string `json:"deceased"`
	DeceasedBoolean     bool   `json:"deceasedBoolean"`
	DeceasedDateTime    string `json:"deceasedDateTime"`
	Gender              string `json:"gender"`
	GeneralPractitioner []struct {
		Display    string `json:"display"`
		ID         string `json:"id"`
		Identifier struct {
		} `json:"identifier"`
		Reference string `json:"reference"`
		Type      string `json:"type"`
	} `json:"generalPractitioner"`
	ID         string `json:"id"`
	Identifier []struct {
		Assigner string `json:"assigner"`
		ID       string `json:"id"`
		Period   struct {
		} `json:"period"`
		System string `json:"system"`
		Type   struct {
		} `json:"type"`
		Use   string `json:"use"`
		Value string `json:"value"`
	} `json:"identifier"`
	ImplicitRules string `json:"implicitRules"`
	Language      string `json:"language"`
	Link          []struct {
		ID    string `json:"id"`
		Other struct {
		} `json:"other"`
		Type string `json:"type"`
	} `json:"link"`
	ManagingOrganization struct {
		Display    string `json:"display"`
		ID         string `json:"id"`
		Identifier struct {
		} `json:"identifier"`
		Reference string `json:"reference"`
		Type      string `json:"type"`
	} `json:"managingOrganization"`
	MaritalStatus struct {
		Coding []struct {
			Code         string `json:"code"`
			Display      string `json:"display"`
			ID           string `json:"id"`
			System       string `json:"system"`
			UserSelected bool   `json:"userSelected"`
			Version      string `json:"version"`
		} `json:"coding"`
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"maritalStatus"`
	Meta struct {
		ID          string `json:"id"`
		LastUpdated string `json:"lastUpdated"`
		Profile     string `json:"profile"`
		Security    []struct {
			Code         string `json:"code"`
			Display      string `json:"display"`
			ID           string `json:"id"`
			System       string `json:"system"`
			UserSelected bool   `json:"userSelected"`
			Version      string `json:"version"`
		} `json:"security"`
		Source string `json:"source"`
		Tag    []struct {
			Code         string `json:"code"`
			Display      string `json:"display"`
			ID           string `json:"id"`
			System       string `json:"system"`
			UserSelected bool   `json:"userSelected"`
			Version      string `json:"version"`
		} `json:"tag"`
		VersionID string `json:"versionId"`
	} `json:"meta"`
	MultipleBirth        string `json:"multipleBirth"`
	MultipleBirthBoolean bool   `json:"multipleBirthBoolean"`
	MultipleBirthInteger int    `json:"multipleBirthInteger"`
	Name                 []struct {
		Family string `json:"family"`
		Given  string `json:"given"`
		ID     string `json:"id"`
		Period struct {
		} `json:"period"`
		Prefix string `json:"prefix"`
		Suffix string `json:"suffix"`
		Text   string `json:"text"`
		Use    string `json:"use"`
	} `json:"name"`
	Photo []struct {
		ContentType string `json:"contentType"`
		Creation    string `json:"creation"`
		Data        string `json:"data"`
		Hash        string `json:"hash"`
		ID          string `json:"id"`
		Language    string `json:"language"`
		Size        int    `json:"size"`
		Title       string `json:"title"`
		URL         string `json:"url"`
	} `json:"photo"`
	Telecom []struct {
		ID     string `json:"id"`
		Period struct {
		} `json:"period"`
		Rank   int    `json:"rank"`
		System string `json:"system"`
		Use    string `json:"use"`
		Value  string `json:"value"`
	} `json:"telecom"`
	Text struct {
		Div    string `json:"div"`
		ID     string `json:"id"`
		Status string `json:"status"`
	} `json:"text"`
}
