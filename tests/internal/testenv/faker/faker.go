package faker

import (
	"fmt"
	"math/rand"
	"strings"
)

type UniqueRand struct {
	generated map[int]bool
}

func (u *UniqueRand) Int() int {
	if u.generated == nil {
		u.generated = make(map[int]bool)
	}

	for {
		i := rand.Int() / 1000000
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

type Faker struct {
	Workspace WorkspaceFaker
	Project   ProjectFaker
	Resource  ResourceFaker
	Dotnet    DotnetFaker
}

func NewFaker() *Faker {
	return &Faker{
		Workspace: WorkspaceFaker{},
		Project:   ProjectFaker{},
		Resource:  ResourceFaker{},
		Dotnet:    DotnetFaker{},
	}
}

type ProjectFaker struct {
}

func (f *ProjectFaker) Name() string {
	list := []string{
		"ShoppingCart",
		"PhotoEditor",
		"PaymentGateway",
		"MachineLearning",
		"AugmentedReality",
		"SocialNetworking",
		"FriendsConnect",
		"BudgetTracker",
		"Telemedicine",
		"EncryptionAlgorithm",
		"RouteFinder",
		"CustomerRelationship",
		"PhysicsEngine",
		"LearningSystem",
		"SmartThermostat",
		"CustomAssistant",
		"VirtualRealitySimulation",
		"CMS",
		"AnalyticsTool",
		"ScalableInfrastructure",
		"CommunityPlatform",
		"VoiceCommand",
		"FitnessApp",
		"InventoryManagement",
		"TravelBooking",
		"ItineraryPlanner",
		"HealthTech",
		"WearableMonitoring",
		"FinTech",
		"P2PLending",
		"EcoFriendlyApp",
		"OnlineCoursePlatform",
		"ChatbotFramework",
		"VirtualWorld",
		"DataInsights",
		"SmartContracts",
		"AerialSurvey",
		"AgriTech",
		"CropMonitoring",
		"ConnectedCar",
		"DonationPlatform",
		"StreamingService",
		"ThreatDetection",
		"Navigation",
		"GPSNavigation",
		"DigitalMarketplace",
		"ContributorRecognition",
		"MedicalRecords",
		"RealEstate",
		"PropertyManagement",
		"SmartAppliances",
		"CryptocurrencyExchange",
		"HealthMonitoring",
		"StorytellingEngine",
		"ProjectManagement",
		"360DegreeExperience",
		"DocumentationPlatform",
		"InfluencerPlatform",
		"CustomProductBuilder",
		"Edutainment",
		"PrecisionFarming",
		"AutonomousVehicles",
		"SelfDrivingCar",
		"AuctionPlatform",
		"WellnessApp",
		"VirtualTours",
	}
	result := list[rand.Intn(len(list))]

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result, rnd.Int())

	return result
}
func (f *ProjectFaker) Path(deep int) string {
	list := []string{
		"/unicorn",
		"/riding",
		"/sunnyday",
		"/cloudsurfing",
		"/treasure/mapquest",
		"/chocolate",
		"/factory/exploration",
		"/time/travel/backtothefuture",
		"/mars",
		"/rover/joyride",
		"/wizardry",
		"/spellcasting/101",
		"/dreamland/sleepwalking",
		"/alien/invasion/observation",
		"/jungle",
		"/explorer/tarzanmode",
		"/galaxy",
		"/farfaraway/odyssey",
		"/robot/danceparty/assemblyline",
		"/zombi",
		"e/apocalypse/survival",
		"/pizza/toppings/customization",
		"/superhero/cape/testing",
		"/dragon/taming/101",
		"/pirate/treasure/hunt",
		"/penguin/snowslide/adventure",
		"/detective/mystery/solvingspree",
		"/moonbase/lunarlanding",
		"/timecapsule",
		"/discovery/centuryoldsnacks",
		"/magiccarpet/ride/flyinghigh",
		"/volcano/lavasurfing",
		"/sasquatch",
		"/sighting/investigation",
		"/candyland/sweettooth/overload",
		"/ghost/haunting",
		"/ghostbustermission",
		"/cyberspace",
		"/virtualreality/escapade",
		"/wizard/elixir/brewing",
		"/chess/grandmaster",
		"/mindbattles",
		"/enchantedforest",
		"/potiongathering",
		"/desert/oasis/treasuremap",
		"/hobbit/secondbreakfast",
		"/adventure",
		"/mirror/dimension/travel",
		"/safari/wildlife",
		"/photography",
		"/robotics/ai/emotionchip",
		"/unicorn/rainbow",
		"/riding",
		"/paralleluniverse",
		"/portalexploration",
		"/mermaid/underwaterconcert",
		"/zoo/animalwhispering",
		"/leprechaun",
		"/gold/chasingrainbows",
		"/steampunk",
		"/timeengine/invention",
		"/music/festival/backsectiontour",
		"/giant/invisiblebeanstalk",
		"/climbing",
		"/fireworks/spectacular",
		"/pyrotechnics",
		"/dinosaur/jurassicadventure",
		"/jedi/lightsaber/training",
		"/fairy/tale/rewriting",
		"/chessboard/queensgambit/masterclass",
		"/astronaut",
		"/moonwalk/danceparty",
	}
	result := ""
	i := 0
	for i < deep {
		items := strings.Split(list[rand.Intn(len(list))], "/")
		for _, path := range items {
			if path != "" {
				if i < deep {
					result += path + "/"
					i++
				} else {
					break
				}
			}
		}
	}

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result[:len(result)-1], rnd.Int())

	return result
}

func (f *ProjectFaker) Group() string {
	list := []string{
		"MobileApp",
		"DataScience",
		"AugmentedReality",
		"Finance",
		"Security",
		"Logistics",
		"Chatbot",
		"CMS",
		"DataVisualization",
		"CloudComputing",
		"VirtualAssistant",
		"Wearables",
		"TravelBooking",
		"Sustainability",
		"EdTech",
		"ArtificialIntelligence",
		"Analytics",
		"Blockchain",
		"Drones",
		"SocialImpact",
		"Entertainment",
		"Cybersecurity",
		"Navigation",
		"IoT",
		"FinTech",
		"GameDevelopment",
		"Collaboration",
		"VirtualReality",
		"SocialNetworking",
		"Education",
		"AgriTech",
		"AutonomousVehicles",
		"ECommerce",
		"HealthTech",
		"RealEstate",
	}
	result := list[rand.Intn(len(list))]

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result, rnd.Int())

	return result
}

func (f *ProjectFaker) Layer() string {
	list := []string{
		"core",
		"core:exception:abstraction",
		"core:exception",
		"presentation",
		"presentation:viewmodel:abstraction",
		"presentation:viewmodel",
		"presentation:view",
		"presentation:controller:abstraction",
		"presentation:controller",
		"communication:request:model:abstractions",
		"communication:request:model",
		"communication:response:model:abstractions",
		"communication:response:model",
		"service",
		"service:contract:abstraction",
		"service:contract",
		"service:abstraction",
		"service:concrete",
		"business:enumeration",
		"business:model:abstraction",
		"business:model",
		"business:validator",
		"business:mapper",
		"persistence:database",
		"persistence:database:entity:abstraction",
		"persistence:database:entity",
		"persistence:database:context",
		"persistence:database:migration",
		"persistence:database:repository",
		"persistence:database:repository:contract",
		"persistence:database:repository:abstraction",
		"persistence:database:repository:concrete",
		"persistence:webservice:model:abstraction",
		"persistence:webservice:model",
		"persistence:webservice:service",
		"messaging:message:abstraction",
		"messaging:message",
		"messaging:messagedispatcher:abstraction",
		"messaging:messagedispatcher",
		"messaging:messagehandler:abstraction",
		"messaging:messagehandler",
		"common:utilities",
		"common:constants",
		"mediator:request:model:abstraction",
		"mediator:request:model",
		"mediator:request:handler:abstraction",
		"mediator:request:handler",
		"mediator:response:model:abstraction",
		"mediator:response:model",
		"mediator:response:handler:abstraction",
		"mediator:response:handler",
		"client-server:rpc:model:abstraction",
		"client-server:rpc:model",
		"client-server:rpc:server",
		"client-server:rpc:client",
		"client-server:http:model:abstraction",
		"client-server:http:model",
		"client-server:http:server",
		"client-server:http:client",
		"client-server:websocket:model:abstraction",
		"client-server:websocket:model",
		"client-server:websocket:server",
		"client-server:websocket:client",
		"client-server:messaging:message:abstraction",
		"client-server:messaging:message",
		"client-server:messaging:publisher",
		"client-server:messaging:subscriber",
	}
	result := list[rand.Intn(len(list))]

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result, rnd.Int())

	return result
}

func (f *ProjectFaker) Package() string {
	list := []string{
		"core",
		"exception",
		"abstraction",
		"exception",
		"presentation",
		"viewmodel",
		"view",
		"controller",
		"communication",
		"request",
		"model",
		"service",
		"contract",
		"concrete",
		"business",
		"enumeration",
		"validator",
		"mapper",
		"database",
		"entity",
		"context",
		"migration",
		"repository",
		"webservice",
		"messaging",
		"messagedispatcher",
		"messagehandler",
		"common",
		"utilities",
		"constants",
	}
	result := list[rand.Intn(len(list))]

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result, rnd.Int())

	return result
}

func (f *ProjectFaker) Set() string {
	list := []string{
		"PatientService",
		"BookingService",
		"ProductService",
		"FinanceService",
		"HRService",
		"PaymentService",
		"NotificationService",
		"MailService",
		"CartService",
		"InventoryService",
		"StockService",
		"AuthService",
		"LogService",
		"ReportService",
		"BillingService",
		"SecurityService",
		"CacheService",
		"DiscountService",
	}
	result := list[rand.Intn(len(list))]

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result, rnd.Int())

	return result
}

type ResourceFaker struct {
}

func (f *ResourceFaker) Name() string {
	list := []string{
		"Customer",
		"Order",
		"Product",
		"ShoppingCart",
		"User",
		"BlogPost",
		"CommentThread",
		"Invoice",
		"ShippingPackage",
		"Reservation",
		"Event",
		"Task",
		"Project",
		"Ticket",
		"Review",
		"Account",
		"Payment",
		"Appointment",
		"TicketBooking",
		"Library",
		"Membership",
		"Wallet",
		"Survey",
		"Poll",
		"Notification",
		"Subscription",
		"Location",
		"Advertisement",
		"EventBooking",
		"EducationCourse",
		"Feedback",
		"Profile",
		"Policy",
		"PolicyHolder",
		"Claim",
		"PolicyPremium",
		"Workflow",
		"LeaveApplication",
		"Document",
		"Calendar",
		"Budget",
		"ProjectTask",
		"Report",
		"Goal",
		"WorkflowInstance",
		"Expense",
		"ReservationBooking",
		"RentalProperty",
		"InventoryItem",
		"FeedbackResponse",
		"RewardPoint",
		"UserPreferences",
		"HealthRecord",
		"Contact",
		"LibraryItem",
		"Checklist",
		"Voucher",
		"Asset",
		"ClaimApproval",
		"MessageThread",
		"ProductReview",
		"Content",
		"Bookmark",
		"SubscriptionPlan",
		"Proposal",
		"Request",
		"Facility",
		"TravelBooking",
		"SecurityRole",
		"AccessControlList",
		"Transaction",
		"LegalCase",
		"Vendor",
		"AdvertisementCampaign",
		"Badge",
		"AppointmentScheduling",
		"BlogCategory",
		"InventoryLocation",
		"ProjectAssignment",
		"TicketResolution",
		"ProductBundle",
		"PatientRecord",
		"AttendanceRecord",
		"RewardRedemption",
		"Complaint",
		"ExpenseReport",
		"CustomerFeedback",
		"OrderCancellation",
		"VenueReservation",
		"ProductRecommendation",
		"EmployeePerformance",
		"RewardCatalog",
		"ProposalApproval",
		"LoanApplication",
		"LibraryMembership",
		"SubscriptionRenewal",
	}
	result := list[rand.Intn(len(list))]

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result, rnd.Int())

	return result
}

type WorkspaceFaker struct {
}

func (f *WorkspaceFaker) Name() string {
	list := []string{
		"PatientService",
		"BookingService",
		"ProductService",
		"FinanceService",
		"HRService",
		"PaymentService",
		"NotificationService",
		"MailService",
		"CartService",
		"InventoryService",
		"StockService",
		"AuthService",
		"LogService",
		"ReportService",
		"BillingService",
		"SecurityService",
		"CacheService",
		"DiscountService",
	}
	result := list[rand.Intn(len(list))]

	rnd := UniqueRand{}
	result = fmt.Sprintf("%v-%v", result, rnd.Int())

	return result
}

type DotnetFaker struct {
}

func (f *DotnetFaker) Package(netVersion string) (string, string) {
	list := make(map[string]string, 0)
	if netVersion == "Net8" {
		list["Microsoft.Extensions.DependencyInjection"] = "8.0.0"
		list["Microsoft.Extensions.Logging"] = "8.0.0"
		list["Microsoft.EntityFrameworkCore.Design"] = "8.0.2"
		list["Microsoft.EntityFrameworkCore.InMemory"] = "8.0.2"
		list["Microsoft.EntityFrameworkCore.Sqlite"] = "8.0.2"
		list["Newtonsoft.Json"] = "13.0.1"
		list["AutoMapper"] = "11.0.0"
		list["FluentValidation"] = "11.1.0"
		list["Moq"] = "4.16.1"
		list["Hangfire"] = "1.7.22"
		list["Serilog"] = "2.10.0"
	}

	var packages []string
	for packageName := range list {
		packages = append(packages, packageName)
	}

	packageName := packages[rand.Intn(len(packages))]
	version := list[packageName]

	return packageName, version
}
