import {
  Button,
  Card,
  CardMedia,
  Rating,
  TextField,
  Typography,
} from "@mui/material";
import Grid from "@mui/material/Unstable_Grid2";
import { Form } from "@remix-run/react";

interface RatingCardProps {
  artistName: string;
  imageUrl: string;
}

export default function RatingCard({ artistName, imageUrl }: RatingCardProps) {
  return (
    <Card sx={{ width: 300 }}>
      <Grid container rowSpacing={1}>
        {imageUrl ? (
          <Grid xs={12} sx={{ height: 300 }}>
            <Grid
              xs={12}
              display="flex"
              justifyContent="center"
              alignItems="center"
              sx={{ maxHeight: 300 }}
            >
              <Typography variant="h5">{artistName}</Typography>
            </Grid>
            <Grid xs={12}>
              <CardMedia
                component="img"
                src={imageUrl}
                alt={`${artistName} image`}
                sx={{ height: 250 }}
              />
            </Grid>
          </Grid>
        ) : (
          <Grid
            xs={12}
            sx={{ height: 300 }}
            display="flex"
            justifyContent="center"
            alignItems="center"
          >
            <Typography variant="h4">{artistName}</Typography>
          </Grid>
        )}
        <Grid xs={12}>
          <Form method="post">
            <input hidden name="artist_name" value={artistName} readOnly />
            <Grid xs={12}>
              <TextField fullWidth name="festival" label="Festival/Concert" />
            </Grid>
            <Grid xs={12}>
              <TextField fullWidth name="year" label="Year" />
            </Grid>
            <Grid xs={12}>
              <Rating precision={0.5} name="rating" />
            </Grid>
            <Grid xs={12}>
              <TextField fullWidth name="comment" label="Comment" />
            </Grid>
            <Grid
              xs={12}
              display="flex"
              justifyContent="center"
              alignItems="center"
            >
              <Button type="submit">Rate</Button>
            </Grid>
          </Form>
        </Grid>
      </Grid>
    </Card>
  );
}
