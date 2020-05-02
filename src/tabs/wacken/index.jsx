import React, { useEffect, useRef, useState } from 'react';
import { Storage } from 'aws-amplify';
import { Grid } from '@material-ui/core';
import EstimationCard from '../../components/estimationCard';
import useRating from '../../hooks/useRating';

const Wacken = () => {
  const wackenBands = useRef([]);
  const [bandsToBeRated, setBandsToBeRated] = useState([]);
  const { data: ratedArtists } = useRating();
  useEffect(() => {
    const getWackenBands = async () => {
      const wackenBandsUrl = await Storage.get('wacken.json');
      const response = await fetch(wackenBandsUrl, { headers: { 'Content-Type': 'application/json' } });
      wackenBands.current = await response.json();
    };
    getWackenBands();
  }, []);
  useEffect(() => {
    if (ratedArtists) {
      const ratedArtistNames = ratedArtists.map((ratedArtist) => ratedArtist.band);
      setBandsToBeRated(wackenBands.current
        .filter((wackenBand) => !ratedArtistNames.includes(wackenBand.artist)));
    }
  }, [ratedArtists]);

  return (
    <Grid container>
      {bandsToBeRated.map((bandToBeRated) => (
        <Grid key={bandToBeRated.artist} item xs={12} sm={12} md={4} lg={3} xl={2}>
          <EstimationCard
            artist={bandToBeRated.artist}
            image={bandToBeRated.image}
            submitCallback={(artist) => {
              setBandsToBeRated(
                bandsToBeRated.filter((band) => band.artist !== artist),
              );
            }}
          />
        </Grid>
      ))}
    </Grid>
  );
};

export default Wacken;
