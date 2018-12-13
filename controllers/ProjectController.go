package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	dao "../dao"
)

// FindAllProjects : for test purpose
func FindAllProjects(w http.ResponseWriter, r *http.Request) {
	projects := dao.QueryProjects()

	type TaskStr struct {
		ID   int    `json:"id"`
		Name string `json:"name"`

		Progress          int      `json:"progress"`
		ProgressByWorklog bool     `json:"progressByWorklog"`
		Relevance         int      `json:"relevance"`
		Type              string   `json:"type"`
		TypeID            string   `json:"typeId"`
		Description       string   `json:"description"`
		Code              string   `json:"code"`
		Level             int      `json:"level"`
		Status            string   `json:"status"`
		Depends           string   `json:"depends"`
		CanWrite          bool     `json:"canWrite"`
		Start             int      `json:"start"`
		Duration          int      `json:"duration"`
		End               int      `json:"end"`
		StartIsMilestone  bool     `json:"startIsMilestone"`
		EndIsMilestone    bool     `json:"endIsMilestone"`
		Collapsed         bool     `json:"collapsed"`
		Assigs            []string `json:"assigs"`
		HasChild          bool     `json:"hasChild"`
	}
	type ResourceStr struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	type RoleStr struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	type GanttData struct {
		Task             []TaskStr     `json:"tasks"`
		Resource         []ResourceStr `json:"resources"`
		Role             []RoleStr     `json:"roles"`
		CanWrite         bool          `json:"canWrite"`
		CanDelete        bool          `json:"canDelete"`
		CanWriteOnParent bool          `json:"canWriteOnParent"`
		CanAdd           bool          `json:"canAdd"`
	}

	var ganttData GanttData
	reg, _ := regexp.Compile("PT_")

	for _, element := range projects {
		isPT := len(reg.FindStringIndex(element.ProjectName)) > 0
		if isPT {
			var task TaskStr
			task.ID = element.ProjectID
			task.Name = element.ProjectName
			task.Status = "STATUS_ACTIVE"
			task.Start = 1396994400000
			task.Duration = 20
			task.End = 1399586399999
			task.Assigs = append(task.Assigs, "a")
			ganttData.Task = append(ganttData.Task, task)
		}

	}
	for i := 0; i < 5; i++ {
		var res ResourceStr
		res.ID = i
		res.Name = "Papiyong"
		ganttData.Resource = append(ganttData.Resource, res)
	}
	for i := 0; i < 5; i++ {
		var role RoleStr
		role.ID = i
		role.Name = "Dev"
		ganttData.Role = append(ganttData.Role, role)
	}
	respondWithJSON(w, http.StatusOK, ganttData)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
