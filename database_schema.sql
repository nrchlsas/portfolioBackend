-- Create Portfolio Database
CREATE DATABASE IF NOT EXISTS portfolio_db;
USE portfolio_db;

-- Create skills table
CREATE TABLE IF NOT EXISTS skills (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create experience table
CREATE TABLE IF NOT EXISTS experience (
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  company VARCHAR(255) NOT NULL,
  period VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  highlights JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create services table
CREATE TABLE IF NOT EXISTS services (
  id INT PRIMARY KEY AUTO_INCREMENT,
  icon VARCHAR(10) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  gradient VARCHAR(100) NOT NULL,
  dark_gradient VARCHAR(100) NOT NULL,
  bg_gradient VARCHAR(100) NOT NULL,
  color VARCHAR(50) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create projects table
CREATE TABLE IF NOT EXISTS projects (
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  tags JSON,
  gradient VARCHAR(100) NOT NULL,
  year INT,
  emoji VARCHAR(10),
  tag_gradients JSON,
  link_color VARCHAR(50),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Create contact table
CREATE TABLE IF NOT EXISTS contact (
  id INT PRIMARY KEY AUTO_INCREMENT,
  email VARCHAR(255) NOT NULL,
  phone VARCHAR(20) NOT NULL,
  location VARCHAR(255) NOT NULL,
  social JSON,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert sample data for skills
INSERT INTO skills (name) VALUES 
  ('React'),
  ('React Native'),
  ('Next.js'),
  ('Vue.js'),
  ('TypeScript'),
  ('JavaScript (ES6+)'),
  ('Redux'),
  ('Zustand'),
  ('Tailwind CSS'),
  ('Styled Components'),
  ('Node.js'),
  ('REST API'),
  ('Java'),
  ('SEEBURGER'),
  ('Oracle Database'),
  ('PL/SQL'),
  ('Docker'),
  ('Linux (Red Hat)'),
  ('Git');

-- Insert sample data for experience
INSERT INTO experience (title, company, period, description, highlights) VALUES 
  (
    'Frontend Engineer — Web & Mobile',
    'PT Ogya Tekno Nusantara',
    '2018 - Present',
    'Developed scalable web and mobile applications for national-level government services (BKKBN). Built frontend systems using React, React Native, and Next.js, including dashboards and mobile apps used by millions of users across Indonesia.',
    '["Built cross-platform applications (web & mobile) with high code reusability", "Implemented offline-first mobile features using AsyncStorage & Background Sync", "Designed responsive dashboards using Recharts, D3.js, and Tailwind CSS", "Optimized performance using Hermes, Flipper, and profiling tools", "Implemented monitoring with Sentry, Firebase Crashlytics, and LogRocket", "Conducted code reviews and mentoring using GitHub PR, ESLint, Prettier"]'
  ),
  (
    'System Integration Engineer',
    'PT Bank Mandiri (SEEBURGER Project)',
    '2018 - Present',
    'Built enterprise integration solutions using Java and SEEBURGER framework for banking systems, ensuring reliable and high-volume data processing.',
    '["Developed middleware integration using SEEBURGER BIS", "Handled large-scale data transactions with Oracle & PL/SQL", "Monitored systems using BIS Monitor and logging tools", "Deployed apps on Red Hat Linux using shell scripting & cron jobs"]'
  );

-- Insert sample data for services
INSERT INTO services (icon, title, description, gradient, dark_gradient, bg_gradient, color) VALUES 
  (
    '💻',
    'Frontend Engineering',
    'Building scalable and high-performance web & mobile applications using React, React Native, Next.js, and Vue.js with modern architecture.',
    'from-indigo-50 to-purple-50',
    'dark:from-indigo-950/30 dark:to-purple-950/30',
    'from-indigo-500 to-purple-500',
    'indigo'
  ),
  (
    '📱',
    'Mobile App Optimization',
    'Optimizing React Native apps for performance, offline-first capabilities, and seamless user experience using modern tools and profiling techniques.',
    'from-emerald-50 to-green-50',
    'dark:from-emerald-950/30 dark:to-green-950/30',
    'from-emerald-500 to-green-500',
    'emerald'
  ),
  (
    '⚙️',
    'System Integration',
    'Developing enterprise-grade integrations using Java, SEEBURGER, and Oracle for high-volume, mission-critical systems.',
    'from-blue-50 to-cyan-50',
    'dark:from-blue-950/30 dark:to-cyan-950/30',
    'from-blue-500 to-cyan-500',
    'blue'
  );

-- Insert sample data for projects
INSERT INTO projects (title, description, tags, gradient, year, emoji, tag_gradients, link_color) VALUES 
  (
    'National Public Service Platform (BKKBN)',
    'Built scalable frontend ecosystem for national services including dashboards and mobile apps serving millions of users across Indonesia.',
    '["React", "React Native", "Next.js"]',
    'from-indigo-500 via-purple-500 to-pink-500',
    2024,
    '🏗️',
    '[{"bg": "bg-indigo-100", "darkBg": "dark:bg-indigo-900/40", "text": "text-indigo-800", "darkText": "dark:text-indigo-300"}, {"bg": "bg-purple-100", "darkBg": "dark:bg-purple-900/40", "text": "text-purple-800", "darkText": "dark:text-purple-300"}, {"bg": "bg-pink-100", "darkBg": "dark:bg-pink-900/40", "text": "text-pink-800", "darkText": "dark:text-pink-300"}]',
    'indigo'
  ),
  (
    'Mobile Offline-First System',
    'Developed mobile apps with offline-first capabilities using AsyncStorage and synchronization strategies for low connectivity environments.',
    '["React Native", "Offline Sync", "Performance"]',
    'from-emerald-500 via-green-500 to-lime-500',
    2023,
    '📱',
    '[{"bg": "bg-emerald-100", "darkBg": "dark:bg-emerald-900/40", "text": "text-emerald-800", "darkText": "dark:text-emerald-300"}, {"bg": "bg-green-100", "darkBg": "dark:bg-green-900/40", "text": "text-green-800", "darkText": "dark:text-green-300"}, {"bg": "bg-lime-100", "darkBg": "dark:bg-lime-900/40", "text": "text-lime-800", "darkText": "dark:text-lime-300"}]',
    'emerald'
  ),
  (
    'Enterprise Integration System',
    'Built middleware integration for banking systems using Java and SEEBURGER, handling high-volume transactions with Oracle database.',
    '["Java", "SEEBURGER", "Oracle"]',
    'from-blue-500 via-cyan-500 to-teal-500',
    2022,
    '⚙️',
    '[{"bg": "bg-blue-100", "darkBg": "dark:bg-blue-900/40", "text": "text-blue-800", "darkText": "dark:text-blue-300"}, {"bg": "bg-cyan-100", "darkBg": "dark:bg-cyan-900/40", "text": "text-cyan-800", "darkText": "dark:text-cyan-300"}, {"bg": "bg-teal-100", "darkBg": "dark:bg-teal-900/40", "text": "text-teal-800", "darkText": "dark:text-teal-300"}]',
    'blue'
  );

-- Insert sample data for contact
INSERT INTO contact (email, phone, location, social) VALUES 
  (
    'nurcholisas123@gmail.com',
    '+62 859-5960-1389',
    'Jakarta, Indonesia',
    '[{"name": "LinkedIn", "url": "https://linkedin.com/in/nurcholis-ahmad-syarif", "initials": "in", "gradient": "from-blue-600 to-blue-700"}, {"name": "GitHub", "initials": "gh", "gradient": "from-gray-800 to-gray-900"}, {"name": "Twitter", "initials": "tw", "gradient": "from-blue-400 to-blue-500"}, {"name": "Instagram", "initials": "ig", "gradient": "from-purple-600 to-pink-600"}]'
  );
