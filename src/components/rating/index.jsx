import React, { useState } from 'react';
import { Rating as RatingMaterialUI } from '@material-ui/lab';
import { Grid, TextField } from '@material-ui/core';
import Button from '@material-ui/core/Button';
import { API } from 'aws-amplify';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import { queryCache, useMutation } from 'react-query';
import useUser from '../../hooks/useUser';

const useStyle = makeStyles({
  root: {
    paddingTop: 20,
  },
});

const addRating = async ({
  artist: band, festival, year, rating, comment, userId, token,
}) => (API.post('musicrating', '/api/v1/ratings/bands', {
  header: { Authorization: `Bearer ${token.data}` },
  body: {
    userId: userId.data, band, festival, year, rating, comment: comment || undefined,
  },
}));

const Rating = ({ bandName }) => {
  const [artist, setArtist] = useState(bandName);
  const [festival, setFestival] = useState('');
  const [year, setYear] = useState('');
  const [rating, setRating] = useState(1);
  const [comment, setComment] = useState('');
  const { userId, token } = useUser();
  const classes = useStyle();
  const [mutate] = useMutation(addRating, {
    onSuccess: () => queryCache.refetchQueries('ratedArtists'),
  });

  const resetRating = () => {
    setArtist('');
    setFestival('');
    setYear('');
    setRating(1);
    setComment('');
  };

  const submitRating = async (event) => {
    event.preventDefault();
    if (userId.data && token.data) {
      if (artist && artist.trim()) {
        await mutate({
          artist, festival, year, rating, comment, userId, token,
        });
        resetRating();
      }
    }
  };

  return (
    <form id={artist ? `rating-form-${artist}` : 'rating-form'} onSubmit={submitRating}>
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
            value={artist}
            onChange={(event) => setArtist(event.target.value)}
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
};

Rating.defaultProps = {
  bandName: '',
};

export default Rating;
