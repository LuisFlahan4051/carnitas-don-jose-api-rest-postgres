package database

import (
	"time"

	"gorm.io/gorm"
)

type ID struct {
	Id         uint      `json:"id"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Deleted_at time.Time `json:"deleted_at"`
}

type Food_types struct {
	ID

	Name  string  `json:"name"`
	Foods []Foods `json:"foods"`
}

type Food_meats struct {
	ID

	Name  string  `json:"name"`
	Foods []Foods `json:"foods"`
}

type Foods struct {
	ID

	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	Food_type_id uint `json:"food_type_id"`
	Food_meat_id uint `json:"food_meat_id"`

	Food_products []Food_products `json:"food_products"`
}

type Drink_sizes struct {
	ID

	Name string `json:"name"`

	Drinks []Drinks `json:"drinks"`
}

type Drink_flavors struct {
	ID

	Name   string   `json:"name"`
	Drinks []Drinks `json:"drinks"`
}

type Drinks struct {
	ID

	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	Drink_size_id   uint `json:"drink_size_id"`
	Drink_flavor_id uint `json:"drink_flavor_id"`

	Drink_products []Drink_products `json:"drink_products"`
}

type Products struct {
	ID

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Photo       string  `json:"photo"`

	Food_products            []Food_products            `json:"food_products"`
	Drink_products           []Drink_products           `json:"drink_products"`
	Branch_products_stock    []Branch_products_stock    `json:"branch_products_stock"`
	Sale_products            []Sale_products            `json:"sale_products"`
	Inventory_products_stock []Inventory_products_stock `json:"inventory_products_stock"`
}

type Food_products struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`

	Food_id    uint `json:"food_id"`
	Product_id uint `json:"product_id"`
}

type Drink_products struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`

	Drink_id   uint `json:"drink_id"`
	Product_id uint `json:"product_id"`
}

// --------
type Supplies struct {
	ID

	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	Branch_supplies_stock    []Branch_supplies_stock    `json:"branch_supplies_stock"`
	Inventory_supplies_stock []Inventory_supplies_stock `json:"inventory_supplies_stock"`
}

type Articles struct {
	ID

	Name        string `json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	Branch_articles_stock    []Branch_articles_stock    `json:"branch_articles_stock"`
	Inventory_articles_stock []Inventory_articles_stock `json:"inventory_articles_stock"`
}

type Safeboxes struct {
	ID

	Cents10 uint ` json:"cents10"`
	Cents50 uint `json:"cents50"`
	Coins1  uint `json:"coins1"`
	Coins2  uint `json:"coins2"`
	Coins5  uint `json:"coins5"`
	Coins10 uint `json:"coins10"`
	Coins20 uint `json:"coins20"`

	Bills20   uint `json:"bills20"`
	Bills50   uint `json:"bills50"`
	Bills100  uint `json:"bills100"`
	Bills200  uint `json:"bills200"`
	Bills500  uint `json:"bills500"`
	Bills1000 uint `json:"bills1000"`

	Branch_safeboxes    []Branch_safeboxes  `json:"branch_safeboxes"`
	Turn_safeboxes      []Turn_safebox      `json:"turn_safeboxes"`
	Inventory_safeboxes []Inventory_safebox `json:"inventory_safeboxes"`
	Action_safebox      `json:"action_safebox"`
}

type Incomes struct {
	ID

	Reason string  `json:"reason"`
	Income float64 `json:"income"`

	Sale_incomes []Sale_incomes `json:"sale_incomes"`
}

type Expenses struct {
	ID

	Reason  string  `json:"reason"`
	Expense float64 `json:"expense"`

	Sale_expenses []Sale_expenses `json:"sale_expenses"`
}

type Arguments struct {
	ID

	Complaint bool   `json:"complaint"`
	Score     uint   `json:"score"`
	Argument  string `json:"argument"`

	Sale_arguments []Sale_arguments `json:"sale_arguments"`
}

//----------------

type Branches struct {
	ID

	Name    string `json:"name"`
	Address string `json:"address"`

	Branch_safeboxes      []Branch_safeboxes      `json:"branch_safeboxes"`
	Branch_products_stock []Branch_products_stock `json:"branch_products_stock"`
	Branch_supplies_stock []Branch_supplies_stock `json:"branch_supplies_stock"`
	Branch_articles_stock []Branch_articles_stock `json:"branch_articles_stock"`
	Users                 []Users                 `json:"users"`
	Branch_user_roles     []Branch_user_roles     `json:"branch_user_roles"`
	Turns                 []Turns                 `json:"turns"`
	Sales                 []Sales                 `json:"sales"`
	Inventories           []Inventories           `json:"inventories"`
	Admin_notifications   []Admin_notifications   `json:"admin_notifications"`
}

type Branch_safeboxes struct {
	ID

	Name    string  `json:"name"`
	Content float64 `json:"content"`

	Branch_id  uint `json:"branch_id"`
	Safebox_id uint `json:"safebox_id"`
}

type Branch_products_stock struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`
	In_use            bool `json:"in_use"`

	Branch_id  uint `json:"branch_id"`
	Product_id uint `json:"product_id"`
}

type Branch_supplies_stock struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`
	In_use            bool `json:"in_use"`

	Branch_id uint `json:"branch_id"`
	Supply_id uint `json:"supply_id"`
}

type Branch_articles_stock struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`
	In_use            bool `json:"in_use"`

	Branch_id  uint `json:"branch_id"`
	Article_id uint `json:"article_id"`
}

//----------------

type Roles struct {
	ID

	Name         string `json:"name"`
	Access_level uint   `json:"access_level"`

	Branch_user_roles  []Branch_user_roles  `json:"branch_user_roles"`
	Turn_user_roles    []Turn_user_roles    `json:"turn_user_roles"`
	Inherit_user_roles []Inherit_user_roles `json:"Inherit_user_roles"`
}

type Users struct {
	ID

	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`

	Admin     bool `json:"admin"`
	Root      bool `json:"root"`
	Verified  bool `json:"verified"`
	Warning   bool `json:"warning"`
	Darktheme bool `json:"darktheme"`

	Active_contract bool `json:"active_contract"`

	Address       string    `json:"address"`
	Born          time.Time `json:"born"`
	Degree_study  string    `json:"degree_study"`
	Relation_ship string    `json:"relation_ship"`
	Curp          string    `json:"curp"`
	Rfc           string    `json:"rfc"`
	Citizen_id    string    `json:"citizen_id"`
	Credential_id string    `json:"credential_id"`
	Origin_state  string    `json:"origin_state"`

	Score     uint   `json:"score"`
	Qualities string `json:"qualities"`
	Defects   string `json:"defects"`

	Branch_id        uint `json:"branch_id"`
	Origin_branch_id uint `json:"origin_branch_id"`

	Inherit_user_roles   []Inherit_user_roles   `json:"Inherit_user_roles"`
	User_phones          []User_phones          `json:"user_phones"`
	User_mails           []User_mails           `json:"user_mails"`
	User_reports         []User_reports         `json:"user_reports"`
	Monetary_bounds      []Monetary_bounds      `json:"monetary_bounds"`
	Monetary_discounts   []Monetary_discounts   `json:"monetary_discounts"`
	Branch_user_roles    []Branch_user_roles    `json:"branch_user_roles"`
	Turns                []Turns                `json:"turns"`
	Turn_user_roles      []Turn_user_roles      `json:"turn_user_roles"`
	Sales                []Sales                `json:"sales"`
	User_safebox_actions []User_safebox_actions `json:"user_safebox_actions"`
	Admin_notifications  []Admin_notifications  `json:"admin_notifications"`
}

type Inherit_user_roles struct {
	gorm.Model

	Role_id uint `json:"role_id"`
	User_id uint `json:"user_id"`
}

type User_phones struct {
	ID

	Phone string `json:"phone"`
	Main  bool   `json:"main"`

	User_id uint `json:"user_id"`
}

type User_mails struct {
	ID

	Mail string `json:"mail"`
	Main bool   `json:"main"`

	User_id uint `json:"user_id"`
}

type User_reports struct {
	ID

	Reason string `json:"reason"`

	User_id uint `json:"user_id"`
}

type Monetary_bounds struct {
	ID

	Reason string  `json:"reason"`
	Bound  float64 `json:"bound"`

	User_id uint `json:"user_id"`
}

type Monetary_discounts struct {
	ID

	Reason   string  `json:"reason"`
	Discount float64 `json:"discount"`

	User_id uint `json:"user_id"`
}

type Branch_user_roles struct {
	ID

	Branch_id uint `json:"branch_id"`
	User_id   uint `json:"user_id"`
	Role_id   uint `json:"role_id"`
}

//--------------------------------

type Turns struct {
	ID

	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
	Active     bool      `json:"active"`

	Income_counter     float64 `json:"income_counter"`
	Netincomes_counter float64 `json:"netincomes_counter"`
	Expenses_counter   float64 `json:"expenses_counter"`

	User_id   uint `json:"user_id"`
	Branch_id uint `json:"branch_id"`

	Turn_user_roles  []Turn_user_roles `json:"turn_user_roles"`
	Sales            []Sales           `json:"sales"`
	Inventories      []Inventories     `json:"inventories"`
	Safebox_received Turn_safebox      `json:"safebox_received"`
	Safebox_finished Turn_safebox      `json:"safebox_finished"`
}

type Turn_user_roles struct {
	gorm.Model

	Login_date  time.Time `json:"login_date"`
	Logout_date time.Time `json:"logout_date"`

	User_id uint `json:"user_id"`
	Role_id uint `json:"role_id"`
	Turn_id uint `json:"turn_id"`
}

type Turn_safebox struct {
	gorm.Model

	Turn_id    uint `gorm:"unique" json:"turn_id"`
	Safebox_id uint `gorm:"unique" json:"safebox_id"`
}

//--------------------------------

type Sales struct {
	ID

	Entry_date   time.Time `json:"entry_date"`
	Exit_date    time.Time `json:"exit_date"`
	Table_number uint      `json:"table_number"`
	Packed       bool      `json:"packed"`

	User_id   uint `json:"user_id"`
	Branch_id uint `json:"branch_id"`
	Turn_id   uint `json:"turn_id"`

	Sale_products  []Sale_products  `json:"sale_products"`
	Sale_incomes   []Sale_incomes   `json:"sale_incomes"`
	Sale_expenses  []Sale_expenses  `json:"sale_expenses"`
	Sale_arguments []Sale_arguments `json:"sale_arguments"`
}

type Sale_products struct {
	ID

	Done               bool    `json:"done"`
	Grammage_quantity  uint    `json:"grammage_quantity"`
	Kilogrammage_price float64 `json:"kilogrammage_price"`
	Unit_quantity      uint    `json:"unit_quantity"`
	Unit_price         float64 `json:"unit_price"`
	Discount           float64 `json:"discount"`
	Tax                float64 `json:"tax"`

	Sale_id    uint `json:"sale_id"`
	Product_id uint `json:"product_id"`
}

type Sale_incomes struct {
	ID

	Sale_id   uint `json:"sale_id"`
	Income_id uint `json:"income_id"`
}

type Sale_expenses struct {
	ID

	Sale_id    uint `json:"sale_id"`
	Expense_id uint `json:"expense_id"`
}

type Sale_arguments struct {
	ID

	Sale_id     uint `json:"sale_id"`
	Argument_id uint `json:"argument_id"`
}

//--------------------------------

type Inventory_types struct {
	ID

	Inventory_type string `json:"inventory_type"`

	Inventories []Inventories `json:"inventories"`
}

type Inventories struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`
	In_use            bool `json:"in_use"`

	Inventory_type_id uint `json:"inventory_type_id"`
	Branch_id         uint `json:"branch_id"`
	Turn_id           uint `json:"turn_id"`

	Inventory_products_stock []Inventory_products_stock `json:"inventory_products_stock"`
	Inventory_supplies_stock []Inventory_supplies_stock `json:"inventory_supplies_stock"`
	Inventory_articles_stock []Inventory_articles_stock `json:"inventory_articles_stock"`
	Inventory_safebox        `json:"inventory_safebox"`
}

type Inventory_products_stock struct {
	ID

	Inventory_id uint `json:"inventory_id"`
	Product_id   uint `json:"product_id"`
}

type Inventory_supplies_stock struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`
	In_use            bool `json:"in_use"`

	Inventory_id uint `json:"inventory_id"`
	Supply_id    uint `json:"supply_id"`
}

type Inventory_articles_stock struct {
	ID

	Unit_quantity     uint `json:"unit_quantity"`
	Grammage_quantity uint `json:"grammage_quantity"`
	In_use            bool `json:"in_use"`

	Inventory_id uint `json:"inventory_id"`
	Article_id   uint `json:"article_id"`
}

type Inventory_safebox struct {
	ID

	Inventory_id uint `json:"inventory_id"`
	Safebox_id   uint `json:"safebox_id"`
}

//--------------------------------

type User_safebox_actions struct {
	ID

	Withdrawal bool `json:"withdrawal"`

	User_id        uint `json:"user_id"`
	Action_safebox `json:"action_safebox"`
}

type Action_safebox struct {
	ID

	Safebox_action_id uint `json:"safebox_action_id"`
	Safebox_id        uint `json:"safebox_id"`
}

type Admin_notifications struct {
	ID

	Type        string `json:"type"`
	Solved      bool   `json:"solved"`
	Description string `json:"description"`

	Branch_id uint                         `json:"branch_id"`
	User_id   uint                         `json:"turn_id"`
	Images    []Admin_notifications_images `json:"images"`
}

type Admin_notifications_images struct {
	ID

	Image string `json:"image"`

	Notification_id uint `json:"notification_id"`
}

type Server_logs struct {
	Id          uint      `json:"id"`
	Create_at   time.Time `json:"create_at"`
	Transaction string    `json:"transaction"`
	User_id     uint      `json:"user_id"`
	Branch_id   uint      `json:"branch_id"`
	Root        bool      `json:"root"`
}
