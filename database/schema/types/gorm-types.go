package database

import (
	"time"

	"gorm.io/gorm"
)

//-------------

type FoodType struct {
	gorm.Model
	Name  string `gorm:"type:varchar(50)" json:"name"`
	Foods []Food `json:"foods"`
}

type FoodMeat struct {
	gorm.Model

	Name  string `gorm:"type:varchar(50)" json:"name"`
	Foods []Food `json:"foods"`
}

type Food struct {
	gorm.Model

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	FoodTypeID uint `json:"food_type_id"`
	FoodMeatID uint `json:"food_meat_id"`

	FoodProducts []FoodProduct `json:"food_products"`
}

type DrinkSize struct {
	gorm.Model

	Name string `gorm:"type:varchar(50)" json:"name"`

	Drinks []Drink `json:"drinks"`
}

type DrinkFlavor struct {
	gorm.Model

	Name   string  `gorm:"type:varchar(50)" json:"name"`
	Drinks []Drink `json:"drinks"`
}

type Drink struct {
	gorm.Model

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	DrinkSizeID   uint `json:"drink_size_id"`
	DrinkFlavorID uint `json:"drink_flavor_id"`

	DrinkProducts []DrinkProduct `json:"drink_products"`
}

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(50)" json:"name"`
	Description string  `json:"description"`
	Price       float64 `gorm:"check: price >= 0" json:"price"`
	Photo       string  `json:"photo"`

	FoodProducts           []FoodProduct           `json:"food_products"`
	DrinkProducts          []DrinkProduct          `json:"drink_products"`
	BranchProductsStock    []BranchProductStock    `json:"branch_products_stock"`
	SalesProducts          []SaleProduct           `json:"sales_products"`
	InventoryProductsStock []InventoryProductStock `json:"inventory_products_stock"`
}

type FoodProduct struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`

	FoodID    uint `json:"food_id"`
	ProductID uint `json:"product_id"`
}

type DrinkProduct struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`

	DrinkID   uint `json:"drink_id"`
	ProductID uint `json:"product_id"`
}

// --------
type Supply struct {
	gorm.Model

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	BranchSuppliesStock    []BranchSupplieStock    `json:"branch_suplies_stock"`
	InventorySuppliesStock []InventorySupplieStock `json:"inventory_supplies_stock"`
}

type Article struct {
	gorm.Model

	Name        string `gorm:"type:varchar(50)" json:"name"`
	Description string `json:"description"`
	Photo       string `json:"photo"`

	BranchArticlesStock    []BranchArticleStock    `json:"branch_articles_stock"`
	InventoryArticlesStock []InventoryArticleStock `json:"inventory_articles_stock"`
}

type Safebox struct {
	gorm.Model

	Cents10 uint `gorm:"check: cents10 >= 0; default: 0" json:"cents10"`
	Cents50 uint `gorm:"check: cents50 >= 0; default: 0" json:"cents50"`
	Coins1  uint `gorm:"check: coins1 >= 0; default: 0" json:"coins1"`
	Coins2  uint `gorm:"check: coins2 >= 0; default: 0" json:"coins2"`
	Coins5  uint `gorm:"check: coins5 >= 0; default: 0" json:"coins5"`
	Coins10 uint `gorm:"check: coins10 >= 0; default: 0" json:"coins10"`
	Coins20 uint `gorm:"check: coins20 >= 0; default: 0" json:"coins20"`

	Bills20   uint `gorm:"check: bills20 >= 0; default: 0" json:"bills20"`
	Bills50   uint `gorm:"check: bills50 >= 0; default: 0" json:"bills50"`
	Bills100  uint `gorm:"check: bills100 >= 0; default: 0" json:"bills100"`
	Bills200  uint `gorm:"check: bills200 >= 0; default: 0" json:"bills200"`
	Bills500  uint `gorm:"check: bills500 >= 0; default: 0" json:"bills500"`
	Bills1000 uint `gorm:"check: bills1000 >= 0; default: 0" json:"bills1000"`

	BranchSafeboxes    []BranchSafebox    `json:"branch_safeboxes"`
	TurnSafeboxes      []TurnSafebox      `json:"turn_safeboxes"`
	InventorySafeboxes []InventorySafebox `json:"inventory_safeboxes"`
	ActionSafebox      `json:"action_safebox"`
}

type Income struct {
	gorm.Model

	Reason      string       `json:"reason"`
	Income      float64      `gorm:"check: income >= 0; not null" json:"income"`
	SaleIncomes []SaleIncome `json:"sale_incomes"`
}

type Expense struct {
	gorm.Model

	Reason       string        `json:"reason"`
	Expense      float64       `gorm:"check: expense >= 0; not null" json:"expense"`
	SaleExpenses []SaleExpense `json:"sale_expenses"`
}

type Argument struct {
	gorm.Model

	Complaint     bool           `json:"complaint"`
	Score         uint           `gorm:"check: score >= 0; default: 0" json:"score"`
	Argument      string         `json:"argument"`
	SaleArguments []SaleArgument `json:"sale_arguments"`
}

//----------------

type Branch struct {
	gorm.Model

	Name    string `gorm:"type:varchar(50)" json:"name"`
	Address string `json:"address"`

	BranchSafeboxes     []BranchSafebox      `json:"branch_safeboxes"`
	BranchProductsStock []BranchProductStock `json:"branch_products_stock"`
	BranchSuppliesStock []BranchSupplieStock `json:"branch_suplies_stock"`
	BranchArticlesStock []BranchArticleStock `json:"branch_articles_stock"`
	Users               []User               `json:"users"`
	BranchUserRoles     []BranchUserRole     `json:"branch_user_roles"`
	Turns               []Turn               `json:"turns"`
	Sales               []Sale               `json:"sales"`
	Inventories         []Inventory          `json:"inventories"`
	AdminNotifications  []AdminNotification  `json:"admin_notifications"`
}

type BranchSafebox struct {
	gorm.Model

	Name string `gorm:"type:varchar(50)" json:"name"`

	BranchID  uint `gorm:"not null" json:"branch_id"`
	SafeboxID uint `gorm:"not null; unique" json:"safebox_id"`
}

type BranchProductStock struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	BranchID  uint `json:"branch_id"`
	ProductID uint `json:"product_id"`
}

type BranchSupplieStock struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	BranchID uint `json:"branch_id"`
	SupplyID uint `json:"supply_id"`
}

type BranchArticleStock struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	BranchID  uint `json:"branch_id"`
	ArticleID uint `json:"article_id"`
}

//----------------

type Role struct {
	gorm.Model

	Name        string `gorm:"type:varchar(50)" json:"name"`
	AccessLevel uint   `gorm:"check: access_level >= 0; default: 0" json:"access_level"`

	BranchUserRoles  []BranchUserRole  `json:"branch_user_roles"`
	TurnUserRoles    []TurnUserRole    `json:"turn_user_roles"`
	InheritUserRoles []InheritUserRole `json:"inherit_user_roles"`
}

type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(50)" json:"name"`
	Lastname string `gorm:"type:varchar(50)" json:"lastname"`
	Username string `gorm:"type:varchar(50)" json:"username"`
	Password string `gorm:"type:varchar(50)" json:"password"`
	Photo    string `json:"photo"`

	Admin     bool `json:"admin"`
	Root      bool `json:"root"`
	Verified  bool `json:"verified"`
	Darktheme bool `json:"darktheme"`

	Active_contract bool `json:"active_contract"`

	Address      string    `json:"address"`
	Born         time.Time `json:"born"`
	DegreeStudy string    `gorm:"type:varchar(50)" json:"degree_study"`
	RelationShip string    `gorm:"type:varchar(50)" json:"relation_ship"`
	Curp         string    `gorm:"type:varchar(50)" json:"curp"`
	Rfc          string    `gorm:"type:varchar(50)" json:"rfc"`
	CitizenID    string    `gorm:"type:varchar(50)" json:"citizen_id"`
	CredentialID string    `gorm:"type:varchar(50)" json:"credential_id"`
	OriginState  string    `gorm:"type:varchar(50)" json:"origin_state"`

	Score          uint   `gorm:"check: score >= 0; default: 0" json:"score"`
	Qualities      string `json:"qualities"`
	Defects        string `json:"defects"`
	OriginBranchID uint   `json:"originBranch_id"`

	BranchID uint `json:"branch_id"`

	InheritUserRoles   []InheritUserRole   `json:"inherit_user_roles"`
	UserPhones         []UserPhone         `json:"user_phones"`
	UserMails          []UserMail          `json:"user_mails"`
	UserReports        []UserReport        `json:"user_reports"`
	MonetaryBounds     []MonetaryBound     `json:"monetary_bounds"`
	MonetaryDiscounts  []MonetaryDiscount  `json:"monetary_discounts"`
	BranchUserRoles    []BranchUserRole    `json:"branch_user_roles"`
	Turns              []Turn              `json:"turns"`
	TurnUserRoles      []TurnUserRole      `json:"turn_user_roles"`
	Sales              []Sale              `json:"sales"`
	UserSafeboxActions []UserSafeboxAction `json:"safebox_actions"`
	AdminNotifications []AdminNotification `json:"notifications"`
}

type InheritUserRole struct {
	gorm.Model

	RoleID uint `json:"role_id"`
	UserID uint `json:"user_id"`
}

type UserPhone struct {
	gorm.Model

	Phone string `gorm:"type:varchar(50)" json:"phone"`
	Main  bool   `json:"main"`

	UserID uint `json:"user_id"`
}

type UserMail struct {
	gorm.Model

	Mail string `gorm:"type:varchar(50)" json:"mail"`
	Main bool   `json:"main"`

	UserID uint `json:"user_id"`
}

type UserReport struct {
	gorm.Model

	Reason string `json:"reason"`

	UserID uint `json:"user_id"`
}

type MonetaryBound struct {
	gorm.Model

	Reason string  `json:"reason"`
	Bound  float64 `json:"bound"`

	UserID uint `json:"user_id"`
}

type MonetaryDiscount struct {
	gorm.Model

	Reason   string  `json:"reason"`
	Discount float64 `json:"discount"`

	UserID uint `json:"user_id"`
}

type BranchUserRole struct {
	gorm.Model

	BranchID uint `json:"branch_id"`
	UserID   uint `json:"user_id"`
	RoleID   uint `json:"role_id"`
}

//--------------------------------

type Turn struct {
	gorm.Model

	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
	Active     bool      `json:"active"`

	UserID        uint           `json:"user_id"`
	BranchID      uint           `json:"branch_id"`
	TurnUserRoles []TurnUserRole `json:"turn_user_roles"`
	TurnSafebox   `json:"turn_safebox"`
	Sales         []Sale      `json:"sales"`
	Inventories   []Inventory `json:"inventories"`
}

type TurnUserRole struct {
	gorm.Model

	Login_date  time.Time `json:"login_date"`
	Logout_date time.Time `json:"logout_date"`

	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
	TurnID uint `json:"turn_id"`
}

type TurnSafebox struct {
	gorm.Model

	TurnID    uint `gorm:"unique" json:"turn_id"`
	SafeboxID uint `gorm:"unique" json:"safebox_id"`
}

//--------------------------------

type Sale struct {
	gorm.Model

	EntryDate    time.Time `json:"entry_date"`
	ExitDate     time.Time `json:"exit_date"`
	Table_number uint      `json:"table_number"`
	Packed       bool      `json:"packed"`

	UserID         uint           `json:"user_id"`
	BranchID       uint           `json:"branch_id"`
	TurnID         uint           `json:"turn_id"`
	SalesProducts  []SaleProduct  `json:"sales_products"`
	SalesIncomes   []SaleIncome   `json:"sales_incomes"`
	SalesExpenses  []SaleExpense  `json:"sales_expenses"`
	SalesArguments []SaleArgument `json:"sales_arguments"`
}

type SaleProduct struct {
	gorm.Model

	Done              bool    `gorm:"default: false" json:"done"`
	GrammageQuantity  uint    `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	KilogrammagePrice float64 `gorm:"check: kilogrammage_price >= 0; default: 0" json:"kilogrammage_price"`
	UnitQuantity      uint    `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	UnitPrice         float64 `gorm:"check: unit_price >= 0; default: 0" json:"unit_price"`
	Discount          float64 `gorm:"check: discount >= 0; default: 0" json:"discount"`
	Tax               float64 `gorm:"check: tax >= 0; default: 0" json:"tax"`

	SaleID    uint `json:"sale_id"`
	ProductID uint `json:"product_id"`
}

type SaleIncome struct {
	gorm.Model

	SaleID   uint `json:"sale_id"`
	IncomeID uint `json:"income_id"`
}

type SaleExpense struct {
	gorm.Model

	SaleID    uint `json:"sale_id"`
	ExpenseID uint `json:"expense_id"`
}

type SaleArgument struct {
	gorm.Model

	SaleID     uint `json:"sale_id"`
	ArgumentID uint `json:"argument_id"`
}

//--------------------------------

type InventoryType struct {
	gorm.Model

	InventoryType string      `json:"inventory_type"`
	Inventories   []Inventory `json:"inventories"`
}

type Inventory struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	InventoryTypeID uint `json:"inventory_type_id"` // INIT/FINAL/LOSSES
	BranchID        uint `json:"branch_id"`
	TurnID          uint `json:"turn_id"`

	InventoryProductsStocks []InventoryProductStock `json:"inventory_products_stocks"`
	InventorySuppliesStocks []InventorySupplieStock `json:"inventory_supplies_stocks"`
	InventoryArticlesStocks []InventoryArticleStock `json:"inventory_articles_stocks"`
	InventorySafebox        `json:"inventory_safebox"`
}

type InventoryProductStock struct {
	gorm.Model

	InventoryID uint `json:"inventory_id"`
	ProductID   uint `json:"product_id"`
}

type InventorySupplieStock struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"inUse"`

	InventoryID uint `json:"inventory_id"`
	SupplyID    uint `json:"supply_id"`
}

type InventoryArticleStock struct {
	gorm.Model

	UnitQuantity     uint `gorm:"check: unit_quantity >= 0; default: 0" json:"unit_quantity"`
	GrammageQuantity uint `gorm:"check: grammage_quantity >= 0; default: 0" json:"grammage_quantity"`
	InUse            bool `json:"in_use"`

	InventoryID uint `json:"inventory_id"`
	ArticleID   uint `json:"article_id"`
}

type InventorySafebox struct {
	gorm.Model

	InventoryID uint `gorm:"unique" json:"inventory_id"`
	SafeboxID   uint `gorm:"unique" json:"safebox_id"`
}

//--------------------------------

type UserSafeboxAction struct {
	gorm.Model

	Withdrawal bool `json:"withdrawal"`

	UserID        uint `json:"user_id"`
	ActionSafebox `json:"action_safebox"`
}

type ActionSafebox struct {
	gorm.Model

	SafeboxActionID uint `gorm:"unique" json:"safebox_action_id"`
	SafeboxID       uint `gorm:"unique" json:"safebox_id"`
}

type AdminNotification struct {
	gorm.Model

	Type        string `json:"type"`
	Solved      bool   `json:"solved"`
	Description string `json:"description"`

	BranchID uint                     `json:"branch_id"`
	UserID   uint                     `json:"turn_id"`
	Images   []AdminNotificationImage `json:"images"`
}

type AdminNotificationImage struct {
	gorm.Model

	Image string `json:"image"`

	NotificationID uint `json:"notification_id"`
}

type ServerLogs struct {
	Id          uint      `json:"id"`
	CreateAt    time.Time `json:"create_at"`
	Transaction string    `json:"transaction"`
	UserID      uint      `json:"user_id"`
	BranchID    uint      `json:"branch_id"`
	Root        bool      `json:"root"`
}
