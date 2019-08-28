import React, { useState } from 'react';
import { Rating as RatingMaterialUI } from '@material-ui/lab';

const Rating = () => {
  const [rating, setRating] = useState(0);
  return (
    <RatingMaterialUI value={rating} onChange={(event, value) => setRating(value)} />
  );
};

export default Rating;
