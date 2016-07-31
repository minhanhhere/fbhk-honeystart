package controller

import (
    "github.com/labstack/echo"
    "fmt"
    "gitlab.com/hs-api-go/database/model"
    "gitlab.com/hs-api-go/manager"
    "net/http"
    "gopkg.in/mgo.v2/bson"
    "gitlab.com/hs-api-go/manager/nganluong"
)

var Transaction struct {
    Prepare echo.HandlerFunc
    Verify  echo.HandlerFunc
}

func init() {
    fmt.Println("init controller.transaction")

    //TODO ADD TOKEN

    Transaction.Prepare = func(c echo.Context) error {
        // PARSE REQUEST
        req := &struct {
            Owner       *model.User `json: "owner"`
            Address     *model.Address `json: "address"`
            PackageId   string `json: "packageId"`
            ProjectId   string `json: "projectId"`
            TotalAmount float64 `json: "totalAmount"`
            Type        string `json: "type"`
        }{}
        if err := c.Bind(req); err != nil {
            return err
        }

        project := manager.ProjectManager.FindById(req.ProjectId)

        // CREATE USER IF NEEDED
        owner := manager.UserManager.GetByEmail(req.Owner.Email)
        if owner == nil {
            fmt.Println("transaction.prepare:", "User not found")
            owner = req.Owner
            owner.Construct()
            owner.Guest = true
            fmt.Println("create new user: ", owner)
        } else if !owner.Guest {
            fmt.Println("transaction.prepare: ", "User found as exist user")
            //TODO: throw error if needed
        } else {
            fmt.Println("transaction.prepare: ", "User found as guest")
        }

        owner.Address = req.Address
        manager.UserManager.Save(owner)

        // CREATE BACKER
        backer := model.NewProjectBacker()
        backer.Owner = owner
        backer.OwnerId = owner.Id
        backer.ProjectId = bson.ObjectIdHex(req.ProjectId)
        backer.PackageId = bson.ObjectIdHex(req.PackageId)
        backer.Address = req.Address
        backer.Type = req.Type
        backer.Quantity = 1
        backer.TotalAmount = req.TotalAmount
        backer.Verified = false
        manager.ProjectManager.SaveBacker(backer)

        buyerInfo := nganluong.BuyerInfo{
            Name: req.Owner.LastName + " " + req.Owner.FirstName,
            Email: req.Owner.Email,
            Phone: req.Owner.Phone,
            Address: req.Address.ToString(),
        }

        transDesc := "Ung ho dam cuoi " + project.Groom.FirstName + " & " + project.Bride.FirstName

        //TODO: change transaction info to something
        url := nganluong.CreateNganLuongPaymentUrl(backer.TotalAmount, backer.Id.Hex(), transDesc, buyerInfo, req.ProjectId, req.PackageId)

        return c.JSON(http.StatusOK, map[string]interface{}{
            "backer": backer,
            "redirect": url,
        })
    }

    Transaction.Verify = func(c echo.Context) error {
        req := &struct {
            TransactionInfo string `json: "transactionInfo"`
            Price           string `json: "price"`
            PaymentId       string `json: "paymentId"`
            PaymentType     string `json: "paymentType"`
            ErrorText       string `json: "errorText"`
            SecureCode      string `json: "secureCode"`
            OrderCode       string `json: "orderCode"`
        }{}
        if err := c.Bind(req); err != nil {
            return err
        }
        fmt.Println("Request:", req)
        valid := nganluong.VerifyNganLuongResult(req.TransactionInfo, req.OrderCode, req.Price, req.PaymentId, req.PaymentType, req.ErrorText, req.SecureCode)
        if valid {
            backer := manager.ProjectManager.FindBackerById(req.OrderCode)
            if backer == nil {
                return echo.NewHTTPError(http.StatusNotFound, "Transaction not found")
            }
            backer.TransactionId = req.PaymentId
            backer.Verified = true
            manager.ProjectManager.SaveBacker(backer)
            return c.JSON(http.StatusOK, backer)
        }
        return echo.NewHTTPError(http.StatusForbidden, "Invalid transaction request")
    }
}

func RegisterTransactionRoute(api *echo.Group) {
    api.POST("/transactions/prepare", Transaction.Prepare)
    api.POST("/transactions/verify", Transaction.Verify)
}