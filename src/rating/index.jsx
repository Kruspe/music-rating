import React, { useState } from 'react';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import { Grid, TextField } from '@material-ui/core';

import './rating.css';


const Rating = () => {
  const [bandName, setBandName] = useState('');
  const [rating, setRating] = useState(0);
  const [comment, setComment] = useState('');
  return (
    <Grid className="rating-container" container justify="center" alignItems="center">
      <Grid item xs={2}>
        <TextField
          id="bandName"
          variant="outlined"
          label="Band Name"
          value={bandName}
          onChange={event => setBandName(event.target.value)}
        />
      </Grid>
      <Grid item className="rating-rating">
        <RatingMaterialUI value={rating} onChange={(event, value) => setRating(value)} />
      </Grid>
      <Grid item xs={5}>
        <TextField
          id="comment"
          fullWidth
          variant="outlined"
          label="Comment"
          value={comment}
          onChange={event => setComment(event.target.value)}
        />
      </Grid>
    </Grid>
  );
};

export default Rating;
