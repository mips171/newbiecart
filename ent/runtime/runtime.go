// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/mikestefanello/pagoda/ent/cartitem"
	"github.com/mikestefanello/pagoda/ent/company"
	"github.com/mikestefanello/pagoda/ent/customer"
	"github.com/mikestefanello/pagoda/ent/order"
	"github.com/mikestefanello/pagoda/ent/orderitem"
	"github.com/mikestefanello/pagoda/ent/passwordtoken"
	"github.com/mikestefanello/pagoda/ent/payment"
	"github.com/mikestefanello/pagoda/ent/product"
	"github.com/mikestefanello/pagoda/ent/productcategory"
	"github.com/mikestefanello/pagoda/ent/schema"
	"github.com/mikestefanello/pagoda/ent/staffmember"
	"github.com/mikestefanello/pagoda/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	cartitemFields := schema.CartItem{}.Fields()
	_ = cartitemFields
	// cartitemDescQuantity is the schema descriptor for quantity field.
	cartitemDescQuantity := cartitemFields[0].Descriptor()
	// cartitem.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	cartitem.QuantityValidator = cartitemDescQuantity.Validators[0].(func(int) error)
	companyFields := schema.Company{}.Fields()
	_ = companyFields
	// companyDescName is the schema descriptor for name field.
	companyDescName := companyFields[0].Descriptor()
	// company.NameValidator is a validator for the "name" field. It is called by the builders before save.
	company.NameValidator = companyDescName.Validators[0].(func(string) error)
	// companyDescBillingContact is the schema descriptor for billing_contact field.
	companyDescBillingContact := companyFields[1].Descriptor()
	// company.BillingContactValidator is a validator for the "billing_contact" field. It is called by the builders before save.
	company.BillingContactValidator = companyDescBillingContact.Validators[0].(func(string) error)
	// companyDescBillingEmail is the schema descriptor for billing_email field.
	companyDescBillingEmail := companyFields[2].Descriptor()
	// company.BillingEmailValidator is a validator for the "billing_email" field. It is called by the builders before save.
	company.BillingEmailValidator = companyDescBillingEmail.Validators[0].(func(string) error)
	// companyDescBillingPhone is the schema descriptor for billing_phone field.
	companyDescBillingPhone := companyFields[3].Descriptor()
	// company.BillingPhoneValidator is a validator for the "billing_phone" field. It is called by the builders before save.
	company.BillingPhoneValidator = companyDescBillingPhone.Validators[0].(func(string) error)
	// companyDescBillingAddress is the schema descriptor for billing_address field.
	companyDescBillingAddress := companyFields[4].Descriptor()
	// company.BillingAddressValidator is a validator for the "billing_address" field. It is called by the builders before save.
	company.BillingAddressValidator = companyDescBillingAddress.Validators[0].(func(string) error)
	customerFields := schema.Customer{}.Fields()
	_ = customerFields
	// customerDescName is the schema descriptor for name field.
	customerDescName := customerFields[0].Descriptor()
	// customer.NameValidator is a validator for the "name" field. It is called by the builders before save.
	customer.NameValidator = customerDescName.Validators[0].(func(string) error)
	// customerDescEmail is the schema descriptor for email field.
	customerDescEmail := customerFields[1].Descriptor()
	// customer.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	customer.EmailValidator = customerDescEmail.Validators[0].(func(string) error)
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescPlacedAt is the schema descriptor for placed_at field.
	orderDescPlacedAt := orderFields[1].Descriptor()
	// order.DefaultPlacedAt holds the default value on creation for the placed_at field.
	order.DefaultPlacedAt = orderDescPlacedAt.Default.(func() time.Time)
	// orderDescBalanceDue is the schema descriptor for balance_due field.
	orderDescBalanceDue := orderFields[2].Descriptor()
	// order.DefaultBalanceDue holds the default value on creation for the balance_due field.
	order.DefaultBalanceDue = orderDescBalanceDue.Default.(float64)
	// order.BalanceDueValidator is a validator for the "balance_due" field. It is called by the builders before save.
	order.BalanceDueValidator = orderDescBalanceDue.Validators[0].(func(float64) error)
	orderitemFields := schema.OrderItem{}.Fields()
	_ = orderitemFields
	// orderitemDescQuantity is the schema descriptor for quantity field.
	orderitemDescQuantity := orderitemFields[0].Descriptor()
	// orderitem.DefaultQuantity holds the default value on creation for the quantity field.
	orderitem.DefaultQuantity = orderitemDescQuantity.Default.(int)
	// orderitem.QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	orderitem.QuantityValidator = orderitemDescQuantity.Validators[0].(func(int) error)
	// orderitemDescUnitPrice is the schema descriptor for unit_price field.
	orderitemDescUnitPrice := orderitemFields[1].Descriptor()
	// orderitem.UnitPriceValidator is a validator for the "unit_price" field. It is called by the builders before save.
	orderitem.UnitPriceValidator = orderitemDescUnitPrice.Validators[0].(func(float64) error)
	passwordtokenFields := schema.PasswordToken{}.Fields()
	_ = passwordtokenFields
	// passwordtokenDescHash is the schema descriptor for hash field.
	passwordtokenDescHash := passwordtokenFields[0].Descriptor()
	// passwordtoken.HashValidator is a validator for the "hash" field. It is called by the builders before save.
	passwordtoken.HashValidator = passwordtokenDescHash.Validators[0].(func(string) error)
	// passwordtokenDescCreatedAt is the schema descriptor for created_at field.
	passwordtokenDescCreatedAt := passwordtokenFields[1].Descriptor()
	// passwordtoken.DefaultCreatedAt holds the default value on creation for the created_at field.
	passwordtoken.DefaultCreatedAt = passwordtokenDescCreatedAt.Default.(func() time.Time)
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescAmount is the schema descriptor for amount field.
	paymentDescAmount := paymentFields[0].Descriptor()
	// payment.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	payment.AmountValidator = paymentDescAmount.Validators[0].(func(float64) error)
	// paymentDescProcessedAt is the schema descriptor for processed_at field.
	paymentDescProcessedAt := paymentFields[4].Descriptor()
	// payment.DefaultProcessedAt holds the default value on creation for the processed_at field.
	payment.DefaultProcessedAt = paymentDescProcessedAt.Default.(func() time.Time)
	productFields := schema.Product{}.Fields()
	_ = productFields
	// productDescName is the schema descriptor for name field.
	productDescName := productFields[0].Descriptor()
	// product.NameValidator is a validator for the "name" field. It is called by the builders before save.
	product.NameValidator = productDescName.Validators[0].(func(string) error)
	// productDescSku is the schema descriptor for sku field.
	productDescSku := productFields[1].Descriptor()
	// product.SkuValidator is a validator for the "sku" field. It is called by the builders before save.
	product.SkuValidator = productDescSku.Validators[0].(func(string) error)
	// productDescDescription is the schema descriptor for description field.
	productDescDescription := productFields[2].Descriptor()
	// product.DefaultDescription holds the default value on creation for the description field.
	product.DefaultDescription = productDescDescription.Default.(string)
	// productDescPrice is the schema descriptor for price field.
	productDescPrice := productFields[3].Descriptor()
	// product.DefaultPrice holds the default value on creation for the price field.
	product.DefaultPrice = productDescPrice.Default.(string)
	// productDescStockCount is the schema descriptor for stock_count field.
	productDescStockCount := productFields[4].Descriptor()
	// product.DefaultStockCount holds the default value on creation for the stock_count field.
	product.DefaultStockCount = productDescStockCount.Default.(int)
	// product.StockCountValidator is a validator for the "stock_count" field. It is called by the builders before save.
	product.StockCountValidator = productDescStockCount.Validators[0].(func(int) error)
	// productDescImageURL is the schema descriptor for image_url field.
	productDescImageURL := productFields[5].Descriptor()
	// product.DefaultImageURL holds the default value on creation for the image_url field.
	product.DefaultImageURL = productDescImageURL.Default.(string)
	// productDescIsActive is the schema descriptor for is_active field.
	productDescIsActive := productFields[6].Descriptor()
	// product.DefaultIsActive holds the default value on creation for the is_active field.
	product.DefaultIsActive = productDescIsActive.Default.(bool)
	// productDescCreatedAt is the schema descriptor for created_at field.
	productDescCreatedAt := productFields[7].Descriptor()
	// product.DefaultCreatedAt holds the default value on creation for the created_at field.
	product.DefaultCreatedAt = productDescCreatedAt.Default.(func() time.Time)
	// productDescUpdatedAt is the schema descriptor for updated_at field.
	productDescUpdatedAt := productFields[8].Descriptor()
	// product.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	product.DefaultUpdatedAt = productDescUpdatedAt.Default.(func() time.Time)
	// product.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	product.UpdateDefaultUpdatedAt = productDescUpdatedAt.UpdateDefault.(func() time.Time)
	productcategoryFields := schema.ProductCategory{}.Fields()
	_ = productcategoryFields
	// productcategoryDescName is the schema descriptor for name field.
	productcategoryDescName := productcategoryFields[0].Descriptor()
	// productcategory.NameValidator is a validator for the "name" field. It is called by the builders before save.
	productcategory.NameValidator = productcategoryDescName.Validators[0].(func(string) error)
	// productcategoryDescDescription is the schema descriptor for description field.
	productcategoryDescDescription := productcategoryFields[1].Descriptor()
	// productcategory.DefaultDescription holds the default value on creation for the description field.
	productcategory.DefaultDescription = productcategoryDescDescription.Default.(string)
	// productcategoryDescImageURL is the schema descriptor for image_url field.
	productcategoryDescImageURL := productcategoryFields[2].Descriptor()
	// productcategory.DefaultImageURL holds the default value on creation for the image_url field.
	productcategory.DefaultImageURL = productcategoryDescImageURL.Default.(string)
	staffmemberFields := schema.StaffMember{}.Fields()
	_ = staffmemberFields
	// staffmemberDescName is the schema descriptor for name field.
	staffmemberDescName := staffmemberFields[0].Descriptor()
	// staffmember.NameValidator is a validator for the "name" field. It is called by the builders before save.
	staffmember.NameValidator = staffmemberDescName.Validators[0].(func(string) error)
	// staffmemberDescEmail is the schema descriptor for email field.
	staffmemberDescEmail := staffmemberFields[1].Descriptor()
	// staffmember.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	staffmember.EmailValidator = staffmemberDescEmail.Validators[0].(func(string) error)
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescNameFirst is the schema descriptor for name_first field.
	userDescNameFirst := userFields[0].Descriptor()
	// user.NameFirstValidator is a validator for the "name_first" field. It is called by the builders before save.
	user.NameFirstValidator = userDescNameFirst.Validators[0].(func(string) error)
	// userDescNameSurname is the schema descriptor for name_surname field.
	userDescNameSurname := userFields[1].Descriptor()
	// user.NameSurnameValidator is a validator for the "name_surname" field. It is called by the builders before save.
	user.NameSurnameValidator = userDescNameSurname.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for phone field.
	userDescPhone := userFields[3].Descriptor()
	// user.PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	user.PhoneValidator = userDescPhone.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[4].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescVerified is the schema descriptor for verified field.
	userDescVerified := userFields[5].Descriptor()
	// user.DefaultVerified holds the default value on creation for the verified field.
	user.DefaultVerified = userDescVerified.Default.(bool)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[6].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
