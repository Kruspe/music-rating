import React, { useState } from 'react';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import { Grid, TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { API, Auth } from 'aws-amplify';
import PropTypes from 'prop-types';

import './rating.css';

const Rating = ({ bandName, onSubmitBehaviour }) => {
  const [band, setBand] = useState(bandName);
  const [festival, setFestival] = useState('');
  const [year, setYear] = useState('');
  const [rating, setRating] = useState(1);
  const [comment, setComment] = useState('');

  const resetRating = () => {
    setBand('');
    setFestival('');
    setYear('');
    setRating(1);
    setComment('');
  };

  const submitRating = async (event) => {
    event.preventDefault();
    if (band && band.trim()) {
      const currentSession = await Auth.currentSession();
      const currentUserInfo = await Auth.currentUserInfo();
      const token = currentSession.getAccessToken().getJwtToken();
      await API.post('musicrating', '/bands', {
        header: { Authorization: `Bearer ${token}` },
        body: {
          user: currentUserInfo.id, band, festival, year, rating, comment: comment || undefined,
        },
      });
      resetRating();
      onSubmitBehaviour();
    }
  };

  return (
    <form id="rating-form" onSubmit={submitRating}>
      <Grid className="rating-container" container justify="center" alignItems="center" spacing={5}>
        <Grid item lg={2} xs={12}>
          <TextField
            id={`${bandName}-band`}
            name="band"
            fullWidth
            required={!bandName}
            InputProps={{ readOnly: !!bandName }}
            variant="outlined"
            label="Band"
            value={band}
            onChange={(event) => setBand(event.target.value)}
          />
        </Grid>
        <Grid item lg={2} xs={12}>
          <TextField
            id={`${bandName}-festival`}
            name="festival"
            fullWidth
            required
            variant="outlined"
            label="Festival"
            value={festival}
            onChange={(event) => setFestival(event.target.value)}
          />
        </Grid>
        <Grid item lg={1} xs={12}>
          <TextField
            id={`${bandName}-year`}
            name="year"
            fullWidth
            required
            variant="outlined"
            label="Year"
            value={year}
            onChange={(event) => setYear(event.target.value)}
          />
        </Grid>
        <Grid item className="rating-rating">
          <RatingMaterialUI name="rating" value={rating} onChange={(event, value) => setRating(value)} />
        </Grid>
        <Grid item lg={4} xs={12}>
          <TextField
            id={`${bandName}-comment`}
            name="comment"
            fullWidth
            variant="outlined"
            label="Comment"
            value={comment}
            onChange={(event) => setComment(event.target.value)}
          />
        </Grid>
        <Grid item lg={1}>
          <Button type="submit" variant="outlined">Submit</Button>
        </Grid>
      </Grid>
    </form>
  );
};

Rating.propTypes = {
  bandName: PropTypes.string,
  onSubmitBehaviour: PropTypes.func,
};

Rating.defaultProps = {
  bandName: '',
  onSubmitBehaviour: () => {},
};

export default Rating;
