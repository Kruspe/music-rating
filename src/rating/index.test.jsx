import React from 'react';
import { fireEvent, render, wait } from '@testing-library/react';
import { API, Auth } from 'aws-amplify';
import Rating from './index';

describe('Rating', () => {
  const isFormInEmptyState = (getByLabelText) => {
    const bandField = getByLabelText(/band \*/i);
    expect(bandField).toHaveValue('');
    expect(bandField).not.toHaveAttribute('readOnly');
    expect(getByLabelText(/festival \*/i)).toHaveValue('');
    expect(getByLabelText(/year \*/i)).toHaveValue('');
    expect(getByLabelText(/1 star/i)).toBeChecked();
    expect(getByLabelText(/comment/i)).toHaveValue('');
  };

  it('should display empty form', () => {
    const { getByLabelText } = render(<Rating />);
    isFormInEmptyState(getByLabelText);
  });

  it('should display form with disabled band field filled in band name', () => {
    const { getByLabelText } = render(<Rating bandName="Bloodbath" />);
    const bandField = getByLabelText(/band/i);
    expect(bandField).toHaveValue('Bloodbath');
    expect(bandField).toHaveAttribute('readOnly');
    expect(bandField).not.toBeRequired();
    expect(getByLabelText(/festival \*/i)).toHaveValue('');
    expect(getByLabelText(/year \*/i)).toHaveValue('');
    expect(getByLabelText(/1 star/i)).toBeChecked();
    expect(getByLabelText(/comment/i)).toHaveValue('');
  });

  describe('submit', () => {
    const postSpy = jest.spyOn(API, 'post').mockImplementation((f) => f);
    beforeEach(() => {
      postSpy.mockClear();
      const currentSessionMock = {
        getAccessToken: () => ({ getJwtToken: () => ('token') }),
      };
      const currentUserInfoMock = { id: 'userId' };
      jest.spyOn(Auth, 'currentSession').mockResolvedValueOnce(currentSessionMock);
      jest.spyOn(Auth, 'currentUserInfo').mockResolvedValueOnce(currentUserInfoMock);
    });
    const fillRatingFields = (getByLabelText) => {
      fireEvent.change(getByLabelText(/band \*/i), { target: { value: 'Bloodbath' } });
      fireEvent.change(getByLabelText(/festival \*/i), { target: { value: 'Wacken' } });
      fireEvent.change(getByLabelText(/year \*/i), { target: { value: '2015' } });
      fireEvent.click(getByLabelText(/5 star/i));
      fireEvent.change(getByLabelText(/comment/i), { target: { value: 'comment' } });
    };

    it('should enter data, save it and call onSubmitBehaviour ', async () => {
      const expectedInit = {
        header: { Authorization: 'Bearer token' },
        body: {
          user: 'userId', band: 'Bloodbath', festival: 'Wacken', year: '2015', rating: 5, comment: 'comment',
        },
      };
      const onSubmitBehaviourMock = jest.fn();

      const {
        getByLabelText, getByText,
      } = render(<Rating onSubmitBehaviour={onSubmitBehaviourMock} />);
      fillRatingFields(getByLabelText);
      fireEvent.submit(getByText(/submit/i));

      await wait(() => expect(postSpy).toHaveBeenCalledTimes(1));
      await wait(() => expect(postSpy).toHaveBeenCalledWith('musicrating', '/bands', expectedInit));
      await wait(() => expect(onSubmitBehaviourMock).toHaveBeenCalledTimes(1));
      isFormInEmptyState(getByLabelText);
    });
    it('should not try to save empty comment', async () => {
      const expectedInit = {
        header: { Authorization: 'Bearer token' },
        body: {
          user: 'userId', band: 'Bloodbath', festival: 'Wacken', year: '2015', rating: 5,
        },
      };

      const { getByLabelText, getByText } = render(<Rating />);
      fillRatingFields(getByLabelText);
      fireEvent.change(getByLabelText(/comment/i), { target: { value: '' } });
      fireEvent.submit(getByText(/submit/i));
      await wait(() => expect(postSpy).toHaveBeenCalledTimes(1));
      await wait(() => expect(postSpy).toHaveBeenCalledWith('musicrating', '/bands', expectedInit));
    });
    it('should require band, festival and year', () => {
      const { getByLabelText, getByText } = render(<Rating />);
      fireEvent.submit(getByText(/submit/i));

      expect(getByLabelText(/band \*/i)).toBeRequired();
      expect(getByLabelText(/festival \*/i)).toBeRequired();
      expect(getByLabelText(/year \*/i)).toBeRequired();
      expect(postSpy).not.toHaveBeenCalled();
    });
  });
});
