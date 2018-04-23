/* eslint-disable import/no-extraneous-dependencies */

import { shallow, mount, createLocalVue } from '@vue/test-utils';
import Vuetify from 'vuetify';
import App from '../../src/App.vue';

const localVue = createLocalVue();

localVue.use(Vuetify);

describe('App.vue', () => {
  it('Should render without crash', () => {
    const wrapper = shallow(App, {
      localVue,
    });
    expect(wrapper).not.toBe(null);
  });

  it('Should render without crash', () => {
    const wrapper = shallow(App, {
      localVue,
    });
    expect(wrapper).toBeInstanceOf(Object);
  });
});
