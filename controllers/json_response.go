package controllers

type JsonResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Total   int         `json:"total"`
	Error   interface{} `json:"error"`
}

func Response(data interface{}, success bool, count int, errors interface{}) JsonResponse {

	response := JsonResponse{}
	response.Success = success
	response.Data = data
	response.Error = errors
	response.Total = count
	return response
}
