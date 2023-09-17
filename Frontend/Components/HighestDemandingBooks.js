// src/components/HighestDemandingBooks.js
import React, { useEffect, useState } from 'react';
import axios from 'axios';

const HighestDemandingBooks = () => {
  const [highestDemandingBooks, setHighestDemandingBooks] = useState([]);

  useEffect(() => {
    // Fetch the highest demanding books from the backend
    fetchHighestDemandingBooks();
  }, []);

  const fetchHighestDemandingBooks = () => {
    // Make a GET request to your Go API to retrieve the highest demanding books
    axios.get('http://your-api-url/highest-demanding-books')
      .then((response) => {
        setHighestDemandingBooks(response.data);
      })
      .catch((error) => {
        console.error('Error fetching highest demanding books:', error);
      });
    
  };

  return (
    <div>
      <h2>Highest Demanding Books</h2>
      <ul>
        {highestDemandingBooks.map((book) => (
          <li key={book.id}>
            {book.title} by {book.author} (Demand: {book.demand})
          </li>
        ))}
      </ul>
    </div>
  );
};

export default HighestDemandingBooks;
