import React from "react";

interface ProjectProps {
  title: string;
  description: string;
  link: string;
}

const ProjectCard: React.FC<ProjectProps> = ({ title, description, link }) => {
  return (
    <div className="p-4 border rounded-lg shadow-md">
      <h2 className="text-xl font-bold">{title}</h2>
      <p className="text-gray-600">{description}</p>
      <a href={link} className="text-blue-500" target="_blank" rel="noopener noreferrer">
        View Project →
      </a>
    </div>
  );
};

export default ProjectCard;
