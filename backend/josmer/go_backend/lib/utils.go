package lib

import "go.mongodb.org/mongo-driver/bson/primitive"

var OmitDatabase = []string{
	"Currency",
	"Country",
	"Email",
	"EmailQueue",
	"EmailTemplate",
	"Language",
	"PasswordRecovery",
	"RootPaymentMethods",
	"RootSiteContent",
	"Site",
	"SiteBuyer",
	"SiteCostTemplate",
	"SitePayments",
	"SitePreference",
	"SocialMedia",
	"State",
	"TermsAndConditions",
	"TimezoneInfo",
	"UploadQueue",
	"User",
	"UserAccess",
	"UserDocument",
	"UserPreference",
	"UserRating",
	"ViaType",
	"HLSFormat",
}

type CheckDomain struct {
	Database    string
	AdminDomain string
	ApiDomain   string
	Exists      bool
	SiteID      primitive.ObjectID
	OwnerID     primitive.ObjectID
}

var RootSites = []string{
	"www.cportatil.com",
	"www.psiconomia.com",
	"www.fattorello.lo",
	"www.cpfback.lo",
}

var RootUser string = "root@consultaportatil.com"
var RootPassword string = "x"
var Hostname string = "cportatil.com"
