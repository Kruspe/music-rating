import React from 'react';
import {
  Table, TableBody, TableCell, TableRow,
} from '@material-ui/core';
import { API, Auth } from 'aws-amplify';

class Overview extends React.Component {
  state = { ratings: undefined };

  async componentDidMount() {
    const currentSession = await Auth.currentSession();
    const currentUserInfo = await Auth.currentUserInfo();
    const token = currentSession.getAccessToken().getJwtToken();

    const ratings = await API.get('musicrating', `/bands/${currentUserInfo.id}`, { header: { Authorization: `Bearer ${token}` } });
    this.setState({ ratings });
  }

  render() {
    const { ratings } = this.state;
    return (
      <React.Fragment>
        {ratings && (
          <Table>
            <TableBody>
              {ratings.map(rating => (
                <TableRow key={rating.band}>
                  <TableCell>{rating.band}</TableCell>
                  <TableCell>{rating.festival}</TableCell>
                  <TableCell>{rating.year}</TableCell>
                  <TableCell>{rating.rating}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        )}
      </React.Fragment>
    );
  }
}

export default Overview;
