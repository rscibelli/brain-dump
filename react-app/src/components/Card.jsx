import React from 'react';

const Card = ({ post }) => {
  const { name, date, title, post: content } = post;

  return (
    <div className="card">
        <div className="card-header bg-secondary text-white">
            <h5 className="card-title mb-0">{title}</h5>
            <div className="d-flex justify-content-between align-items-center">
            <h6 className="card-subtitle mb-1">{name}</h6>
            <span className="card-date">{date}</span>
            </div>
        </div>
        <div className="card-body text-start">
            <p className="card-text">{content}</p>
        </div>
    </div>
  );
};

export default Card;