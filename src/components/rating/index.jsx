import React, { useState, useContext } from 'react';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import { Grid, TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { API } from 'aws-amplify';
import PropTypes from 'prop-types';
import UserContext from '../../context/UserContext';
import { makeStyles } from '@material-ui/core/styles';

const useStyle = makeStyles({
  root: {
    paddingTop: 20,
},
});

const Rating = ({ bandName, onSubmitBehaviour }) => {
  const classes = useStyle();
  const [band, setBand] = useState(bandName);
  const [festival, setFestival] = useState('');
  const [year, setYear] = useState('');
  const [rating, setRating] = useState(1);
  const [comment, setComment] = useState('');
  const { userId, jwtToken } = useContext(UserContext);

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
      await API.post('musicrating', '/api/v1/ratings/bands', {
        header: { Authorization: `Bearer ${jwtToken}` },
        body: {
          user: userId, band, festival, year, rating, comment: comment || undefined,
        },
      });
      resetRating();
      onSubmitBehaviour();
    }
  };

  return (
    <form id={band ? `rating-form-${band}` : 'rating-form'} onSubmit={submitRating}>
      <Grid className={classes.root} container justify="center" alignItems="center" spacing={5}>
        <Grid item xs={12}>
          <TextField
            id={bandName ? `${bandName}-band` : 'band'}
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
        <Grid item xs={12}>
          <TextField
            id={bandName ? `${bandName}-festival` : 'festival'}
            name="festival"
            fullWidth
            required
            variant="outlined"
            label="Festival"
            value={festival}
            onChange={(event) => setFestival(event.target.value)}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            id={bandName ? `${bandName}-year` : 'year'}
            name="year"
            fullWidth
            required
            variant="outlined"
            label="Year"
            value={year}
            onChange={(event) => setYear(event.target.value)}
          />
        </Grid>
        <Grid item>
          <RatingMaterialUI
            name={bandName ? `${bandName}-rating` : 'rating'}
            value={rating}
            onChange={(event, value) => setRating(value)}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            id={bandName ? `${bandName}-comment` : 'comment'}
            name="comment"
            fullWidth
            variant="outlined"
            label="Comment"
            value={comment}
            onChange={(event) => setComment(event.target.value)}
          />
        </Grid>
        <Grid item>
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
