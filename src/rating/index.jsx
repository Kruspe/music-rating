import React, { useState } from 'react';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import { Grid, TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { API, Auth } from 'aws-amplify';

import './rating.css';

const Rating = () => {
  const [bandName, setBandName] = useState('');
  const [rating, setRating] = useState(1);
  const [comment, setComment] = useState('');

  const submitRating = async (event) => {
    event.preventDefault();
    if (bandName && bandName.trim()) {
      const currentSession = await Auth.currentSession();
      const token = currentSession.getAccessToken().getJwtToken();
      await API.post('musicrating', '/bands', {
        header: { Authorization: `Bearer ${token}` },
        body: { bandName, rating, comment },
      });
    }
  };

  return (
    <form id="rating-form" onSubmit={submitRating}>
      <Grid className="rating-container" container justify="center" alignItems="center" spacing={5}>
        <Grid item xs={2}>
          <TextField
            id="bandName"
            variant="outlined"
            label="Band Name"
            value={bandName}
            onChange={event => (setBandName(event.target.value))}
          />
        </Grid>
        <Grid item className="rating-rating">
          <RatingMaterialUI name="rating" value={rating} onChange={(event, value) => setRating(value)} />
        </Grid>
        <Grid item xs={4}>
          <TextField
            id="comment"
            fullWidth
            variant="outlined"
            label="Comment"
            value={comment}
            onChange={event => setComment(event.target.value)}
          />
        </Grid>
        <Grid item xs={1}>
          <Button type="submit" variant="outlined"> Submit </Button>
        </Grid>
      </Grid>
    </form>
  );
};

export default Rating;
