package controllers

import(
    "fmt"
    "github.com/gin-gonic/gin"
    "database/sql"
)

type Orders struct {
    orderId int
    orderedAt string
    customerName string
    items Items
}

type Items struct {
    itemId int
    itemCode string
    description string
    quantity int
}

var (
    db *sql.DB
    err error
)

func CreateOrder(ctx *gin.Context) {
    var order = Orders{}

    sqlStatement := `
    INSERT INTO orders (order_id, customer_name, ordered_at)
    VALUES ($1, $2, $3, $4)
    RETURNING *
    `

    err = db.QueryRow(sqlStatement, 1, "Tom Jerry", "2022-09-27").Scan(&order.orderId, &order.customerName, &order.orderedAt)

    if err != nil {
        panic(err)
    }

    fmt.Printf("New Item Data : %+v\n", order)
}

func GetOrder(ctx *gin.Context) {
    var results = []Orders{}

    sqlStatement := `SELECT item_id, item_code from items`

    rows, err := db.Query(sqlStatement);

    if err != nil {
        panic(err)
    }
    defer rows.Close()

    for rows.Next() {
        var item = Orders{}
        err = rows.Scan(&item.items.itemId, &item.items.itemCode)

        if err != nil {
            panic(err)
        }

        results = append(results, item)
    }

    fmt.Println("Items:", results)

}

func UpdateOrder(ctx *gin.Context) {
    // TO DO
}

func DeleteOrder(ctx *gin.Context) {
    // TO DO
}