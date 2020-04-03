import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { makeStyles } from '@material-ui/core/styles';
import Rating from '../rating';

const useStyle = makeStyles({
  button: {
    color: '#FFFFFF',
    background: '#000000',
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    width: '100%',
    height: 360,
    cursor: 'pointer',
  },
  image: {
    maxWidth: 300,
    maxHeight: 300,
  },
});

const EstimationCard = (props) => {
  const classes = useStyle();
  const [isRatingVisible, setIsRatingVisible] = useState(false);
  const { artist, image, submitCallback } = props;

  return (
    <>
      <button className={classes.button} type="button" onClick={() => setIsRatingVisible(!isRatingVisible)}>
        <h1>{artist}</h1>
        {image && <img className={classes.image} src={image} alt={artist} />}
      </button>
      {isRatingVisible && (
      <Rating
        bandName={artist}
        onSubmitBehaviour={() => submitCallback(artist)}
      />
      )}
    </>
  );
};

EstimationCard.propTypes = {
  artist: PropTypes.string.isRequired,
  image: PropTypes.string,
  submitCallback: PropTypes.func.isRequired,
};

EstimationCard.defaultProps = {
  image: undefined,
};

export default EstimationCard;
