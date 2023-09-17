// BookDetails.js
import React, { useState } from 'react';

const BookDetails = ({ book }) => {
  const [quantity, setQuantity] = useState(1);

  // Implement handling quantity change and adding to cart

  return (
    <div>
      <h2>{book.title}</h2>
      {/* Display book details */}
      <input
        type="number"
        value={quantity}
        onChange={(e) => setQuantity(e.target.value)}
      />
      <button onClick={() => addToCart(book.id, quantity)}>Add to Cart</button>
    </div>
  );
};
