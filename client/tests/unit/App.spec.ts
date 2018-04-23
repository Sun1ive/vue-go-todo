/* eslint-disable import/no-extraneous-dependencies */

import { shallow, mount, createLocalVue } from '@vue/test-utils';
import sinon from 'sinon';
import Vuetify from 'vuetify';
import Vue from 'vue';
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
  it('Should contain correct wrapper', () => {
    const wrapper = shallow(App, {
      localVue,
    });
    expect(wrapper).toBeInstanceOf(Object);
  });
  it('Should be vue instance', () => {
    const wrapper = shallow(App, {
      localVue,
    });
    expect(wrapper.isVueInstance()).toBe(true);
  });
  it('Should contain correct classes', () => {
    const wrapper = shallow(App, {
      localVue,
    });
    expect(wrapper.classes()).toContain('application');
    expect(wrapper.classes()).toContain('theme--light');
  });
  it('Should set correct data', () => {
    const todos = ['hello', 'world'];
    const wrapper = mount(App, {
      localVue,
    });
    wrapper.setData({ todos });
    // @ts-ignore
    expect(wrapper.vm.todos).toBe(todos);
  });
  it('Should set correct methods', () => {
    const clickHandler = sinon.stub();
    const wrapper = mount(App, {
      localVue,
      propsData: { clickHandler },
    });
    const divArray = wrapper.findAll('div');
    expect(clickHandler.called).toBe(false);
    divArray.trigger('click');
    expect(clickHandler.called).toBe(true);
  });
});
