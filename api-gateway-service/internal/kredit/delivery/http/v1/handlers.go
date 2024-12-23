package v1

import (
	"fmt"
	"net/http"

	"github.com/0x5w4/kredit-plus/api-gateway-service/config"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/dto"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/commands"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/mapper"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/queries"
	"github.com/0x5w4/kredit-plus/api-gateway-service/internal/kredit/service"
	httpErrors "github.com/0x5w4/kredit-plus/pkg/http_errors"
	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/go-playground/validator"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type kreditHandler struct {
	group  *echo.Group
	logger *logger.AppLogger
	cfg    *config.Config
	ps     *service.KreditService
	v      *validator.Validate
}

func NewKreditHandler(
	group *echo.Group,
	logger *logger.AppLogger,
	cfg *config.Config,
	ps *service.KreditService,
	v *validator.Validate,
) *kreditHandler {
	return &kreditHandler{group: group, logger: logger, cfg: cfg, ps: ps, v: v}
}

// CreateKonsumen
// @Tags Kredit
// @Summary Create konsumen
// @Description Create new konsumen item
// @Accept json
// @Produce json
// @Param request body dto.CreateKonsumenHttpRequest true "Request"
// @Success 201 {object} dto.CreateKonsumenResponseDto
// @Router /konsumen [post]
func (h *kreditHandler) CreateKonsumen() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		httpRequest := &dto.CreateKonsumenHttpRequest{}
		if err := c.Bind(httpRequest); err != nil {
			h.logger.SLogger.Warn("Bind", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		createKonsumenRequest, err := mapper.CreateKonsumenHttpToDto(*httpRequest)
		if err != nil {
			h.logger.SLogger.Warn("mapper.CreateKonsumenHttpToDto", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		createKonsumenRequest.IdKonsumen, err = uuid.NewV7()
		if err != nil {
			h.logger.SLogger.Warn("uuid.NewV7", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		fmt.Println(createKonsumenRequest)
		if err := h.v.StructCtx(ctx, createKonsumenRequest); err != nil {
			h.logger.SLogger.Warn("validate", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		fmt.Println("4---------------------------")
		if err := h.ps.Commands.CreateKonsumen.Handle(ctx, commands.NewCreateKonsumenCommand(&createKonsumenRequest)); err != nil {
			h.logger.SLogger.Warn("CreateKonsumen", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusCreated, dto.CreateKonsumenResponseDto{IdKonsumen: createKonsumenRequest.IdKonsumen})
	}
}

// CreateLimit
// @Tags Kredit
// @Summary Create limit
// @Description Create new limit item
// @Accept json
// @Produce json
// @Param request body dto.CreateLimitHttpRequest true "Request"
// @Success 201 {object} dto.CreateLimitResponseDto
// @Router /limit [post]
func (h *kreditHandler) CreateLimit() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		httpRequest := &dto.CreateLimitHttpRequest{}
		if err := c.Bind(httpRequest); err != nil {
			h.logger.SLogger.Warn("Bind", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		createLimitRequest, err := mapper.CreateLimitHttpToDto(*httpRequest)
		if err != nil {
			fmt.Println(err)
			h.logger.SLogger.Warn("mapper.CreateLimitHttpToDto", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		createLimitRequest.IdLimit, err = uuid.NewV7()
		if err != nil {
			h.logger.SLogger.Warn("uuid.NewV7", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		fmt.Println(createLimitRequest)
		if err := h.v.StructCtx(ctx, createLimitRequest); err != nil {
			h.logger.SLogger.Warn("validate", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		fmt.Println("4---------------------------")
		if err := h.ps.Commands.CreateLimit.Handle(ctx, commands.NewCreateLimitCommand(&createLimitRequest)); err != nil {
			h.logger.SLogger.Warn("CreateLimit", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusCreated, dto.CreateLimitResponseDto{IdLimit: createLimitRequest.IdLimit})
	}
}

// CreateTransaksi
// @Tags Kredit
// @Summary Create transaksi
// @Description Create new transaksi item
// @Accept json
// @Produce json
// @Param request body dto.CreateTransaksiHttpRequest true "Request"
// @Success 201 {object} dto.CreateTransaksiResponseDto
// @Router /transaksi [post]
func (h *kreditHandler) CreateTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		httpRequest := &dto.CreateTransaksiHttpRequest{}
		if err := c.Bind(httpRequest); err != nil {
			h.logger.SLogger.Warn("Bind", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		createTransaksiRequest, err := mapper.CreateTransaksiHttpToDto(*httpRequest)
		if err != nil {
			h.logger.SLogger.Warn("mapper.CreateKonsumenHttpToDto", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		createTransaksiRequest.IdTransaksi, err = uuid.NewV7()
		if err != nil {
			h.logger.SLogger.Warn("uuid.NewV7", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		fmt.Println(createTransaksiRequest)
		if err := h.v.StructCtx(ctx, createTransaksiRequest); err != nil {
			h.logger.SLogger.Warn("validate", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		fmt.Println("4---------------------------")
		if err := h.ps.Commands.CreateTransaksi.Handle(ctx, commands.NewCreateTransaksiCommand(&createTransaksiRequest)); err != nil {
			h.logger.SLogger.Warn("CreateTransaksi", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusCreated, dto.CreateTransaksiResponseDto{IdTransaksi: createTransaksiRequest.IdTransaksi})
	}
}

// GetLimit
// @Tags Kredit
// @Summary Get limit
// @Description Get limit by id
// @Accept json
// @Produce json
// @Param id_limit path string true "Id Limit"
// @Param id_konsumen path string true "Id Konsumen"
// @Success 200 {object} dto.GetLimitResponseDto
// @Router /limit/{id} [get]
func (h *kreditHandler) GetLimit() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		idLimit, err := uuid.Parse(c.Param("id_limit"))
		if err != nil {
			h.logger.SLogger.Warn("uuid.Parse", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}
		idKonsumen, err := uuid.Parse(c.Param("id_konsumen"))
		if err != nil {
			h.logger.SLogger.Warn("uuid.Parse", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		query := queries.NewGetLimitQuery(idLimit, idKonsumen)
		response, err := h.ps.Queries.GetLimit.Handle(ctx, query)
		if err != nil {
			h.logger.SLogger.Warn("Queries.GetLimit", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusOK, response)
	}
}

// GetTransaksi
// @Tags Kredit
// @Summary Get transaksi
// @Description Get transaksi by id
// @Accept json
// @Produce json
// @Param id_transaksi path string true "Id Transaksi"
// @Param id_konsumen path string true "Id Konsumen"
// @Success 200 {object} dto.GetTransaksiResponseDto
// @Router /transaksi/{id} [get]
func (h *kreditHandler) GetTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		idTransaksi, err := uuid.Parse(c.Param("id_transaksi"))
		if err != nil {
			h.logger.SLogger.Warn("uuid.Parse", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}
		idKonsumen, err := uuid.Parse(c.Param("id_konsumen"))
		if err != nil {
			h.logger.SLogger.Warn("uuid.Parse", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		query := queries.NewGetTransaksiQuery(idTransaksi, idKonsumen)
		response, err := h.ps.Queries.GetTransaksi.Handle(ctx, query)
		if err != nil {
			h.logger.SLogger.Warn("Queries.GetTransaksi", err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		return c.JSON(http.StatusOK, response)
	}
}
