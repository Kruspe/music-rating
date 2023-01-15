import PropTypes from 'prop-types';
import { withAuthenticationRequired } from '@auth0/auth0-react';

export default function AuthenticationGuard({ component }) {
  const Component = withAuthenticationRequired(component);
  return <Component />;
}

AuthenticationGuard.propTypes = {
  component: PropTypes.func.isRequired,
};
