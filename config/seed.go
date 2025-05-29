package config

import (
	"encoding/json"
	"log"
	"wannn-site-rebuild-api/models"
)

func SeedDatabase() {
	// Seed Experiences
	experiences := []models.Experience{
		{
			Title:   "New Venture & Technology Incubation Intern",
			Company: "PT. XL Axiata Tbk",
			Period:  "Sep 2024 - Dec 2024",
			Description: mustJSON([]string{
				"Led development of Roadinspex, a road damage detection system, implementing 15+ RESTful APIs with Sequelize ORM and JWT authentication",
				"Designed comprehensive UML diagrams and conducted thorough testing (API, Unit, Integration, Functional) with detailed documentation",
				"Optimized HelloMet safety monitoring system by migrating to Jetson Nano with SSD MobileNet, achieving 7x performance improvement",
				"Developed WANalyze public transport dashboard using Home Assistant with Frigate for real-time occupancy monitoring",
				"Contributed to Smart AC Automation project for Indomaret using Thingsboard, integrating IoT devices for temperature control",
			}),
		},
		{
			Title:   "Data Labeler",
			Company: "Retrux Studio",
			Period:  "Nov 2024 - Dec 2024",
			Description: mustJSON([]string{
				"Labeled and validated 200+ supermarket shelf images for stock availability detection, ensuring high-quality training data",
				"Utilized labelImg for precise bounding box annotation and collaborated with a 5-member team on 1,000+ image dataset",
				"Conducted peer reviews of annotations to maintain dataset accuracy and consistency",
				"Enhanced ML development efficiency by providing clean, validated datasets that reduced validation workload",
			}),
		},
		{
			Title:   "Computer Vision",
			Company: "Barunastra ITS RoboBoat Team",
			Period:  "Jan 2023 - Dec 2024",
			Description: mustJSON([]string{
				"Developed vision-side pipeline for Autonomous Surface Vehicles (ASV), including object detection, tracking, and counting using YOLOv5",
				"Implemented 2-step detection feature that improved buoy detection accuracy by 90% through color-based recognition",
				"Optimized computer vision processing by migrating to edge devices, reducing power consumption by 46%",
				"Managed all computer-related systems for extended ASV deployments, ensuring reliable operation",
			}),
		},
	}

	// Seed Projects
	projects := []models.Project{
		{
			Title:       "Roadinspex",
			Description: "A road damage detection system built with Node.js, Express, PostgreSQL, and Sequelize. This project is a part of my internship at PT. XL Axiata Tbk. I was responsible for developing the backend of the system, including the RESTful APIs and the database schema.",
			Technologies: mustJSON([]string{
				"Node.js",
				"Express",
				"PostgreSQL",
				"Sequelize",
				"JWT",
			}),
			Link: "https://roadinspex.xdevelopment.my.id/",
		},
		{
			Title:       "MIoT (Multimedia and Internet of Things) Laboratorium Website",
			Description: "A profile website of Multimedia and Internet of Things (MIoT) Laboratorium at Computer Engineering Department, Institut Teknologi Sepuluh Nopember. This project is our responsibility as Web Development Team in MIoT Laboratorium. I was responsible for developing the 10+ reusable components and 2 key pages, including the 'Practicums' page and the 'Our Researchs' page.",
			Technologies: mustJSON([]string{
				"React",
				"Tailwind CSS",
				"JavaScript",
				"Framer Motion",
				"React Router",
				"React Icons",
			}),
			Link: "https://miot-lab.vercel.app/",
		},
		{
			Title:       "Soil Monitoring Website",
			Description: "A soil monitoring website built with Vite, React.js, Tailwind CSS, Express.js, and InfluxDB. This project is a part of my freelance as Web Developer. The key features are the real-time soil moisture (Nitrogen, pH, Phosphorus, Potassium) monitoring and the dashboard interface.",
			Technologies: mustJSON([]string{
				"Vite",
				"React.js",
				"Tailwind CSS",
				"Express.js",
				"InfluxDB",
			}),
			Link: "https://soilmonitor.my.id/",
		},
		{
			Title:       "Water Level Monitoring Website",
			Description: "A water level monitoring website built with Vite, Vue.js, Tailwind CSS, Express.js, and MongoDB. This project is a part of my freelance as Web Developer. The key features are the real-time water level monitoring, toggling, and automating the water pump.",
			Technologies: mustJSON([]string{
				"Vite",
				"React.js",
				"Tailwind CSS",
				"Express.js",
				"MongoDB",
			}),
			Link: "https://watermonitor.site/",
		},
		{
			Title:       "YOLOv5-ROS2",
			Description: "A ROS2 Humble Hawksbill package for object detection using YOLOv5. This project is a part of my job as a Computer Vision at Barunastra ITS RoboBoat Team. I was responsible for developing vision-side pipeline, including object detection, object tracking, and object counting utilizing YOLOv5 model.",
			Technologies: mustJSON([]string{
				"Python",
				"ROS2",
				"YOLOv5",
				"OpenCV",
				"PyTorch",
			}),
			Link: "https://github.com/wannn-one/yolov5-ros2",
		},
	}

	// Seed Skill Categories
	skillCategories := []models.SkillCategory{
		{
			Title: "Languages",
			Skills: mustJSON([]string{
				"C++",
				"Python",
				"JavaScript",
				"Go",
				"SQL",
				"HTML/CSS",
				"Shell Script",
			}),
		},
		{
			Title: "Frameworks & Libraries",
			Skills: mustJSON([]string{
				"React.js",
				"Vue.js",
				"Tailwind CSS",
				"Express.js",
				"Fiber",
				"GORM",
				"Sequelize",
				"JWT",
				"ROS2",
				"PyTorch",
				"OpenCV",
			}),
		},
		{
			Title: "DevOps & Tools",
			Skills: mustJSON([]string{
				"Docker",
				"Git",
				"GitHub",
				"CI/CD",
				"Linux",
				"AWS",
				"Visual Studio Code",
				"Nginx",
				"Apache2",
				"Postman",
			}),
		},
	}

	// Insert seed data
	for _, exp := range experiences {
		result := DB.Create(&exp)
		if result.Error != nil {
			log.Printf("Error seeding experience: %v", result.Error)
		}
	}

	for _, proj := range projects {
		result := DB.Create(&proj)
		if result.Error != nil {
			log.Printf("Error seeding project: %v", result.Error)
		}
	}

	for _, cat := range skillCategories {
		result := DB.Create(&cat)
		if result.Error != nil {
			log.Printf("Error seeding skill category: %v", result.Error)
		}
	}

	log.Println("Database seeding completed")
}

// Helper function to convert slice to JSON string
func mustJSON(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
} 