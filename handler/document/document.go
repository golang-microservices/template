package document

import (
	"bytes"
	"io"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"github.com/golang-common-packages/cloud-storage"

	"github.com/golang-common-packages/template/config"
	"github.com/golang-common-packages/template/model"
)

// Handler manage all request and dependency
type Handler struct {
	*config.Environment
}

// New return a new Handler
func New(env *config.Environment) *Handler {
	return &Handler{env}
}

// Handler function will register all path to echo.Echo
func (h *Handler) Handler(e *echo.Group) {
	e.GET("/document", h.list(), h.JWT.Middleware(h.Config.Token.Accesstoken.PublicKey), h.Cache.Middleware(h.Hash), h.Monitor.Middleware())
	e.POST("/document", h.save(), h.JWT.Middleware(h.Config.Token.Accesstoken.PublicKey), h.Cache.Middleware(h.Hash), h.Monitor.Middleware())
	// e.GET("/drive", h.files(), h.JWT.Middleware(h.Config.Token.Accesstoken.PublicKey), h.Monitor.Middleware())
	// e.POST("/drive", h.upload(), h.JWT.Middleware(h.Config.Token.Accesstoken.PublicKey), h.Monitor.Middleware())
	// e.DELETE("/drive", h.delete(), h.JWT.Middleware(h.Config.Token.Accesstoken.PublicKey), h.Monitor.Middleware())
	// e.GET("/drive/donwload", h.donwload(), h.JWT.Middleware(h.Config.Token.Accesstoken.PublicKey), h.Monitor.Middleware())
}

// localhost:3000/api/v1/document?limit=3
// localhost:3000/api/v1/document?limit=3&lastid=5cee0e7af554bfbe838882c2
func (h *Handler) list() echo.HandlerFunc {
	return func(c echo.Context) error {
		results, err := h.Database.GetALL(h.Config.Service.Database.MongoDB.DB, h.Config.Service.Database.Collection.Document, c.QueryParam("lastid"), c.QueryParam("limit"), reflect.TypeOf(model.Document{}))
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		documents, ok := results.(*[]model.Document)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError, "Can not bind the result to model")
		}
		//fmt.Println((*documents)[0])

		return c.JSON(http.StatusOK, documents)
	}
}

func (h *Handler) save() echo.HandlerFunc {
	return func(c echo.Context) error {
		validate := validator.New()
		request := model.Document{}

		// Bind request body to struct
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Validate request body struct
		if err := validate.Struct(request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		result, err := h.Database.Create(h.Config.Service.Database.MongoDB.DB, h.Config.Service.Database.Collection.Document, request)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (h *Handler) files() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileInfo := &cloudStorage.FileModel{Path: ""}
		files, err := h.Storage.List(fileInfo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, files)
	}
}

func (h *Handler) upload() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		mimeType := c.FormValue("mimeType")
		parentID := c.FormValue("parentID")

		file, err := c.FormFile("file")
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		src, err := file.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		defer src.Close()

		buf := bytes.NewBuffer(nil)
		if _, err := io.Copy(buf, src); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		// This will return a value of type bytes.Reader which implements the io.Reader (and io.ReadSeeker) interface.
		// Don't worry about them not being the same "type". io.Reader is an interface and can be implemented by many different types.
		data := bytes.NewReader(buf.Bytes())

		fileInfo := &cloudStorage.FileModel{
			Name:     name,
			MimeType: mimeType,
			ParentID: parentID,
			Content:  data,
		}

		result, err := h.Storage.Upload(fileInfo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (h *Handler) delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileInfo := &cloudStorage.FileModel{
			SourcesID: c.QueryParam("fileid"),
		}

		err := h.Storage.Delete(fileInfo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusOK)
	}
}

func (h *Handler) donwload() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileInfo := &cloudStorage.FileModel{
			SourcesID: c.QueryParam("fileid"),
		}

		files, err := h.Storage.Download(fileInfo)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.File(files.(string))
	}
}
