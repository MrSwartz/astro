package handler

import (
	"astro"
	"astro/pkg/tools"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) getPicOfTheDay(w http.ResponseWriter, r *http.Request) {
	ur, params := tools.UrlConstructor("https://api.nasa.gov/planetary/apod", r)
	result := tools.MakeRequests(ur)

	if result[ur].Error != nil {
		tools.SendResponse(w, astro.Response{
			Status: http.StatusInternalServerError,
			Error:  astro.Error{Message: result[ur].Error.Error()},
		})
		return
	}

	pic := astro.Picture{}
	if err := json.Unmarshal(result[ur].Body, &pic); err != nil {
		tools.SendResponse(w, astro.Response{
			Status: http.StatusInternalServerError,
			Error:  astro.Error{Message: err.Error()},
		})
		return
	}

	url := pic.Url
	if params["hd"] == "true" {
		url = pic.Hdurl
	}

	result2 := tools.MakeRequests(url)
	if result2[url].Error != nil {
		tools.SendResponse(w, astro.Response{
			Status: http.StatusInternalServerError,
			Error:  astro.Error{Message: result2[url].Error.Error()},
		})
	}

	pic.BinaryPic = string(result2[url].Body)

	if params["store"] == "true" {
		if n, err := h.services.Insert(pic); err != nil && n != 1 {
			tools.SendResponse(w, astro.Response{
				Status: http.StatusInternalServerError,
				Error:  astro.Error{Message: fmt.Sprintf("rows affected %d/1, error: %q", n, err)},
			})
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "image/jpg")
	w.Write(result2[url].Body)

}

func (h *Handler) getPicturesFromStorage(w http.ResponseWriter, r *http.Request) {
	_, params := tools.UrlConstructor("", r)

	sd, errSd := tools.ParseTime(params["start_date"])
	ed, errEd := tools.ParseTime(params["end_date"])
	date, errD := tools.ParseTime(params["date"])

	var pics []astro.Picture
	var err error
	var status = http.StatusOK
	if errSd == nil && errEd == nil {
		pics, err = h.services.GetByDateRange(sd, ed)
	} else if errD == nil {
		pics, err = h.services.GetByDate(date)
	} else {
		status = http.StatusInternalServerError
		err = fmt.Errorf("wrong url param combination or err while parsing: %q", err)
	}

	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}

	tools.SendResponse(w, astro.Response{
		Status:   status,
		Error:    astro.Error{Message: errMsg},
		Pictures: pics,
	})
}
