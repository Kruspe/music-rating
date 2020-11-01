import { API } from 'aws-amplify';
import {
  fireEvent, render, screen, waitFor,
} from '@testing-library/react';
import Rating from './index';
import useUser from '../../hooks/useUser';
import useRating from '../../hooks/useRating';

jest.mock('../../hooks/useUser');

const renderWithUseRating = () => {
  const Wrapper = () => {
    useRating();
    return (<Rating />);
  };
  return render(<Wrapper />);
};

describe('Rating', () => {
  const expectFormToBeEmpty = () => {
    const bandField = screen.getByLabelText(/band \*/i);
    expect(bandField).toHaveValue('');
    expect(bandField).not.toHaveAttribute('readOnly');
    expect(screen.getByLabelText(/festival \*/i)).toHaveValue('');
    expect(screen.getByLabelText(/year \*/i)).toHaveValue('');
    expect(screen.getByLabelText(/1 star/i)).toBeChecked();
    expect(screen.getByLabelText(/comment/i)).toHaveValue('');
  };

  beforeEach(() => {
    useUser.mockReturnValue({ userId: { data: 'userId' } });
  });

  it('should display empty form', () => {
    render(<Rating />);
    expectFormToBeEmpty();
  });

  it('should prefill artist field and make it readonly', () => {
    render(<Rating bandName="Bloodbath" />);
    const bandField = screen.getByLabelText(/band/i);
    expect(bandField).toHaveValue('Bloodbath');
    expect(bandField).toHaveAttribute('readOnly');
    expect(bandField).not.toBeRequired();
    expect(screen.getByLabelText(/festival \*/i)).toHaveValue('');
    expect(screen.getByLabelText(/year \*/i)).toHaveValue('');
    expect(screen.getByLabelText(/1 star/i)).toBeChecked();
    expect(screen.getByLabelText(/comment/i)).toHaveValue('');
  });

  describe('submit', () => {
    const fillRatingFields = () => {
      fireEvent.change(screen.getByLabelText(/band \*/i), { target: { value: 'Bloodbath' } });
      fireEvent.change(screen.getByLabelText(/festival \*/i), { target: { value: 'Wacken' } });
      fireEvent.change(screen.getByLabelText(/year \*/i), { target: { value: '2015' } });
      fireEvent.click(screen.getByLabelText(/5 star/i));
      fireEvent.change(screen.getByLabelText(/comment/i), { target: { value: 'comment' } });
    };

    it('should not post rating when userId is undefined', () => {
      useUser.mockReturnValue({ userId: { data: undefined } });
      render(<Rating />);
      fillRatingFields();
      fireEvent.submit(screen.getByText(/submit/i));

      expect(API.post).not.toHaveBeenCalled();
      expect(screen.getByLabelText(/band \*/i)).toHaveValue('Bloodbath');
      expect(screen.getByLabelText(/festival \*/i)).toHaveValue('Wacken');
      expect(screen.getByLabelText(/year \*/i)).toHaveValue('2015');
      expect(screen.getByLabelText(/5 star/i)).toBeChecked();
      expect(screen.getByLabelText(/comment/i)).toHaveValue('comment');
    });

    it('should submit rating and fetch ratings again', async () => {
      const expectedInit = {
        header: { Authorization: 'Bearer token' },
        body: {
          userId: 'userId', artist: 'Bloodbath', festival: 'Wacken', year: '2015', rating: 5, comment: 'comment',
        },
      };
      renderWithUseRating();
      await waitFor(() => expect(API.get).toHaveBeenCalledTimes(1));
      fillRatingFields();
      fireEvent.submit(screen.getByText(/submit/i));
      await waitFor(() => expect(API.post).toHaveBeenCalledWith('musicrating', '/api/v1/ratings/bands', expectedInit));
      expect(API.post).toHaveBeenCalledTimes(1);
      await waitFor(() => expect(API.get).toHaveBeenCalledTimes(2));
      expectFormToBeEmpty();
    });

    it('should submit rating without comment', async () => {
      const expectedInit = {
        header: { Authorization: 'Bearer token' },
        body: {
          userId: 'userId', artist: 'Bloodbath', festival: 'Wacken', year: '2015', rating: 5,
        },
      };
      renderWithUseRating();
      await waitFor(() => expect(API.get).toHaveBeenCalledTimes(1));
      fillRatingFields();
      fireEvent.change(screen.getByLabelText(/comment/i), { target: { value: '' } });
      fireEvent.submit(screen.getByText(/submit/i));
      await waitFor(() => expect(API.post).toHaveBeenCalledWith('musicrating', '/api/v1/ratings/bands', expectedInit));
      expect(API.post).toHaveBeenCalledTimes(1);
      await waitFor(() => expect(API.get).toHaveBeenCalledTimes(2));
      expectFormToBeEmpty();
    });
    it('should require artist, festival and year', () => {
      render(<Rating />);
      fireEvent.submit(screen.getByText(/submit/i));

      expect(screen.getByLabelText(/band \*/i)).toBeRequired();
      expect(screen.getByLabelText(/festival \*/i)).toBeRequired();
      expect(screen.getByLabelText(/year \*/i)).toBeRequired();
      expect(API.post).not.toHaveBeenCalled();
    });
  });
});
