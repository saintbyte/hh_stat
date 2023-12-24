package structures

type Negotiations struct {
	Items []struct {
		ID    string `json:"id"`
		State struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"state"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
		Resume    struct {
			ID           string `json:"id"`
			Title        string `json:"title"`
			URL          string `json:"url"`
			AlternateURL string `json:"alternate_url"`
		} `json:"resume"`
		ViewedByOpponent bool   `json:"viewed_by_opponent"`
		HasUpdates       bool   `json:"has_updates"`
		MessagesURL      string `json:"messages_url"`
		URL              string `json:"url"`
		Counters         struct {
			Messages       int `json:"messages"`
			UnreadMessages int `json:"unread_messages"`
		} `json:"counters"`
		ChatStates struct {
			ResponseReminderState struct {
				Allowed bool `json:"allowed"`
			} `json:"response_reminder_state"`
		} `json:"chat_states"`
		Source          string `json:"source"`
		ChatID          int64  `json:"chat_id"`
		MessagingStatus string `json:"messaging_status"`
		DeclineAllowed  bool   `json:"decline_allowed"`
		Read            bool   `json:"read"`
		HasNewMessages  bool   `json:"has_new_messages"`
		Hidden          bool   `json:"hidden"`
		Vacancy         struct {
			ID                     string `json:"id"`
			Premium                bool   `json:"premium"`
			Name                   string `json:"name"`
			Department             any    `json:"department"`
			HasTest                bool   `json:"has_test"`
			ResponseLetterRequired bool   `json:"response_letter_required"`
			Area                   struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"area"`
			Salary struct {
				From     int    `json:"from"`
				To       any    `json:"to"`
				Currency string `json:"currency"`
				Gross    bool   `json:"gross"`
			} `json:"salary"`
			Type struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"type"`
			Address           any    `json:"address"`
			ResponseURL       any    `json:"response_url"`
			SortPointDistance any    `json:"sort_point_distance"`
			PublishedAt       string `json:"published_at"`
			CreatedAt         string `json:"created_at"`
			Archived          bool   `json:"archived"`
			ApplyAlternateURL string `json:"apply_alternate_url"`
			ShowLogoInSearch  any    `json:"show_logo_in_search"`
			InsiderInterview  any    `json:"insider_interview"`
			URL               string `json:"url"`
			AlternateURL      string `json:"alternate_url"`
			Employer          struct {
				ID           string `json:"id"`
				Name         string `json:"name"`
				URL          string `json:"url"`
				AlternateURL string `json:"alternate_url"`
				LogoUrls     struct {
					Num90    string `json:"90"`
					Num240   string `json:"240"`
					Original string `json:"original"`
				} `json:"logo_urls"`
				VacanciesURL         string `json:"vacancies_url"`
				AccreditedItEmployer bool   `json:"accredited_it_employer"`
				Trusted              bool   `json:"trusted"`
			} `json:"employer"`
			ProfessionalRoles []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"professional_roles"`
		} `json:"vacancy,omitempty"`
		Vacancy0 struct {
			ID                     string `json:"id"`
			Premium                bool   `json:"premium"`
			Name                   string `json:"name"`
			Department             any    `json:"department"`
			HasTest                bool   `json:"has_test"`
			ResponseLetterRequired bool   `json:"response_letter_required"`
			Area                   struct {
				ID   string `json:"id"`
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"area"`
			Salary struct {
				From     int    `json:"from"`
				To       any    `json:"to"`
				Currency string `json:"currency"`
				Gross    bool   `json:"gross"`
			} `json:"salary"`
			Type struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"type"`
			Address           any    `json:"address"`
			ResponseURL       any    `json:"response_url"`
			SortPointDistance any    `json:"sort_point_distance"`
			PublishedAt       string `json:"published_at"`
			CreatedAt         string `json:"created_at"`
			Archived          bool   `json:"archived"`
			ApplyAlternateURL string `json:"apply_alternate_url"`
			InsiderInterview  any    `json:"insider_interview"`
			URL               string `json:"url"`
			AlternateURL      string `json:"alternate_url"`
			Employer          struct {
				ID                   string `json:"id"`
				Name                 string `json:"name"`
				URL                  string `json:"url"`
				AlternateURL         string `json:"alternate_url"`
				LogoUrls             any    `json:"logo_urls"`
				VacanciesURL         string `json:"vacancies_url"`
				AccreditedItEmployer bool   `json:"accredited_it_employer"`
				Trusted              bool   `json:"trusted"`
			} `json:"employer"`
			ProfessionalRoles []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"professional_roles"`
		} `json:"vacancy,omitempty"`
	} `json:"items"`
	Found   int `json:"found"`
	Pages   int `json:"pages"`
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
