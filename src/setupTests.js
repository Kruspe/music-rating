import '@testing-library/jest-dom/extend-expect';
import { queryCache } from 'react-query';

beforeEach(() => {
  queryCache.clear();
});