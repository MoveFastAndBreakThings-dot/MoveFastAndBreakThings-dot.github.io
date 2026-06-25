package content

type Profile struct {
	Name     string
	Role     string
	Location string
	Email    string
	Bio      []string
}

type Job struct {
	Title   string
	Company string
	Period  string
	Bullets []string
	Tags    []string
}

type Project struct {
	Name        string
	Event       string
	Description string
	Tags        []string
	URL         string
}

type SkillGroup struct {
	Category string
	Items    []string
}

type Link struct {
	Label string
	URL   string
}

var MyProfile = Profile{
	Name:     "Samardeep Singh",
	Role:     "Engineering Student @ University of Alberta",
	Location: "Edmonton, AB, Canada",
	Email:    "samarde1@ualberta.ca",
	Bio: []string{
		"Engineering student at University of Alberta. Passionate about using ML in interdisciplinary areas like computational biology & drug design, physics, telecommunication, sustainability, renewable energy, hardware systems, etc. I mean I find it very innovative when 2 things from different fields get combined to form or predict a newly solution that nobody had thought of, kinda like Alphafold. ",		"Previously 4 months as a Data Analyst / ML Intern at Molecule AI — drug-target interaction analysis, molecular toxicity classifiers, SMILES preprocessing pipelines.",
		"Software team member at UAARG working on Shepard — the onboard imaging and autopilot system running on Raspberry Pi. Building real-time ArUco marker detection pipelines with OpenCV and streaming annotated video output to Emu, UAARG's ground station frontend, in support of competition missions at Unmanned Systems Canada and AUVSI SUAS.",
	},
}

var Jobs = []Job{
	{
		Title:   "Data Analyst / Machine Learning Intern",
		Company: "Molecule AI Pvt. Ltd. — New Delhi, India",
		Period:  "Jan 2024 – Apr 2024",
		Bullets: []string{
			"Engineered end-to-end preprocessing pipelines for molecular property datasets (SMILES strings, bioactivity labels) enabling reproducible feature engineering for GNN and Ridge Regression models.",
			"Conducted EDA on drug-target interaction datasets; statistical modeling identified physicochemical feature correlations that improved virtual screening hit rates.",
			"Trained and benchmarked scikit-learn classifiers for molecular toxicity prediction, reducing false-positive rates in compound filtering and accelerating lead prioritization workflows.",
		},
		Tags: []string{"PyTorch", "GNN", "Scikit-learn", "Python", "pandas", "RDKit", "Drug Discovery"},
	},
	{
		Title:   "Asset Protection Associate",
		Company: "Walmart Canada — Edmonton, AB",
		Period:  "Oct 2025 – Present",
		Bullets: []string{
			"Execute systematic inventory audits and data-driven loss investigations using internal reporting tools.",
			"Produce structured incident reports and root-cause analyses that identify anomalous patterns and inform process improvements.",
			"Coordinate with department managers to implement corrective action plans reducing documented inventory discrepancies.",
		},
		Tags: []string{"Data Analysis", "Reporting", "Operations"},
	},
	{
		Title:   "Software Team Member – Imaging & Onboard Perception",
		Company: "UAARG (University of Alberta Aerial Robotics Group)",
		Period:  "Oct 2025 – Present",
		Bullets: []string{
			"Contributing to Shepard, UAARG's onboard drone software stack running on Raspberry Pi — implementing real-time ArUco marker detection with OpenCV for precise localization and target identification during autonomous flight missions.",
			"Developing a live video streaming pipeline that overlays ArUco bounding box annotations and telemetry onto frames transmitted to Emu, UAARG's ground station imaging frontend, enabling operators to monitor detection confidence in real time.",
			"Supporting UAARG's annual entries in the Unmanned Systems Canada Student UAS Competition and AUVSI SUAS, where the full imaging pipeline — from onboard capture to ground-station analysis — is evaluated under competition conditions.",
			"Exploring the integration of deep reinforcement learning to replace hand-tuned flight controllers with learned autonomous navigation policies — designing initial reward structures and simulation environments for UAV decision-making.",
		},
		Tags: []string{"Python", "OpenCV", "ArUco", "Raspberry Pi", "Computer Vision", "Video Streaming", "UAV"},
	},
}

var Projects = []Project{
	{
		Name:        "TorsionNet — Deep RL for Molecular Conformer Generation -- In Progress",
		Event:       "HackED 2026",
		Description: "Replicated TorsionNet (NeurIPS), a deep RL approach to conformer search. Built a custom OpenAI Gym environment with RDKit, trained a PPO agent with an MPNN backbone and curriculum learning, deployed inference via FastAPI and Docker. Achieved 90% reduction in generation time vs. brute-force search.",
		Tags:        []string{"PyTorch", "PyTorch Geometric", "RDKit", "PPO", "GNN", "FastAPI", "Docker"},
		URL:         "",
	},
	{
		Name:        "CoSound — Mindful Social Listening",
		Event:       "NatHacks 2024 · Honorary Mention",
		Description: "Real-time ambient-audio-driven music recommendation system, physically installed at Cameron Library, University of Alberta. Ridge Regression classifier on ESC-50 dataset; 5-dimensional user preference embeddings from Librosa audio features matched against song profiles via live NFC-tag voting.",
		Tags:        []string{"Python", "Scikit-learn", "Librosa", "React", "Node.js", "PostgreSQL", "Docker"},
		URL:         "",
	},
	{
		Name:        "MACE-MP-0 ML Force Field — Browser-Native Deployment via ONNX",
		Event:       "Side Project",
		Description: "Converted MACE-MP-0 (3.8M parameter ML force field) from PyTorch to ONNX with dynamic graph support, enabling browser-native inference via ONNX Runtime Web — no Python, no server. Built an ASE-compatible Calculator wrapping the ONNX model; energy predictions match PyTorch to float32 precision (Δ < 2×10⁻⁶ eV) across 6 molecules from water to benzene. Deployed an interactive Marimo notebook via WebAssembly where Python drives live ML inference through a Python-to-JS bridge — state-of-the-art ML force fields with zero installation.",
		Tags:        []string{"Python", "PyTorch", "ONNX", "ONNX Runtime Web", "ASE", "WebAssembly", "Marimo", "Computational Chemistry"},
		URL:         "",
	},
	{
		Name:        "SSH Terminal Portfolio",
		Event:       "Side Project",
		Description: "Interactive portfolio served over raw SSH — no browser, no login required. Built a Go SSH server using Wish middleware that spawns a full Bubbletea TUI per session. Features 5 navigable tabs, 3 switchable color themes, true-color portrait rendering via half-block Unicode (▄) with 24-bit ANSI, OSC 8 clickable hyperlinks, and automated CI/CD to Fly.io via GitHub Actions. Packaged as a ~7 MB Alpine Docker image. Connect: ssh -p 23234 samar-personal-portfolio.fly.dev",
		Tags:        []string{"Go", "Wish", "Bubbletea", "Lipgloss", "Docker", "Fly.io", "SSH", "GitHub Actions", "CI/CD"},
		URL:         "",
	},
}

var SkillGroups = []SkillGroup{
	{
		Category: "ML / AI",
		Items:    []string{"PyTorch", "PyTorch Geometric", "Scikit-learn", "Deep Learning", "PPO / RL", "GNN", "Statistical Modeling", "NumPy", "SciPy", "Librosa"},
	},
	{
		Category: "Data Science",
		Items:    []string{"Python", "pandas", "SQL", "PostgreSQL", "Feature Engineering", "EDA", "Jupyter", "matplotlib", "seaborn"},
	},
	{
		Category: "Robotics / CV",
		Items:    []string{"OpenCV", "ArUco", "YOLO", "Raspberry Pi", "Video Streaming", "UAV Imaging", "Computer Vision Pipelines", "Unity Simulations"},
	},
	{
		Category: "Engineering / MLOps",
		Items:    []string{"Docker", "Git", "SSH", "Linux", "FastAPI", "REST APIs", "CI/CD", "AWS", "Supabase", "React + typescript framework", "GoLang", "MLflow", "ONNX", "WebAssembly"},
	},
}

var Links = []Link{
	{Label: "GitHub", URL: "https://github.com/MoveFastAndBreakThings-dot?tab=repositories"},
	{Label: "LinkedIn", URL: "https://www.linkedin.com/in/samardeep-singh-488b4024b/"},
	{Label: "Email", URL: "mailto:samarde1@ualberta.ca"},
	{Label: "Resume", URL: "https://drive.google.com/file/d/1bvnNffGLOVFWwAoVa1sZzd6BVQ-xCNEX/view?usp=sharing"},
}

var Extracurriculars = []string{
	"Runner-Up — Enbridge Case Competition, University of Alberta (2025)",
	"Runner-Up — Global Health Case Competition, MSF / Hep-C in Cox's Bazar (2026)",
	"Participant — HackED 2026, University of Alberta Engineering Students' Society",
	"Honorary Mention — NatHacks 2024, NeurAlbertaTech Hacks",
}
