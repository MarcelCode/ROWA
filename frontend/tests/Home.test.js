import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import vuetify from "vuetify"
import Vuex, { mapState } from 'vuex'
import HomeTopRow from "../src/components/main/HomeTopRow"
import FarmInfo from "../src/components/home/FarmInfo";
import CatTree from "../src/components/home/Farm/CatTree";
import Home from '../src/views/Home.vue'
import axios from 'axios';

jest.mock('axios');
describe('Home',  () => {
  let wrapper;
  let store;
  let localVue
  let data;
  beforeEach(() => {
    localVue =  createLocalVue()
    localVue.use(vuetify)
    localVue.use(Vuex)
    store = new Vuex.Store({
      state: {
        farm_active: true
      }
    })
    axios.get.mockResolvedValue({})
    wrapper = shallowMount(Home, {store, localVue})
  })

  it('has a created hook', async()  => {
    expect(wrapper.isVueInstance()).toBe(true)  
  })

  it('contains the right components', () => {
    expect(wrapper.contains(FarmInfo)).toBe(true)
    expect(wrapper.contains(HomeTopRow)).toBe(true)
    expect(wrapper.contains(CatTree)).toBe(true)
  })

  it('returns the right sensor data', async()=>{
    data = {
      data: [{
      datetime: null,
      temperature: 22,
      light_intensity: 150
    }]
    }
    axios.get.mockResolvedValue(data)
    wrapper = shallowMount(Home, {store, localVue})
    await wrapper.vm.$nextTick()
    expect(wrapper.vm.sensor_data.temperature).toEqual(22)
    expect(wrapper.vm.sensor_data.light_intensity).toEqual(150)
  })

  it('sets the right interval', async() => {
    data = {
      data: [{
      datetime: null,
      temperature: 22,
      light_intensity: 150
    }]
    }
    axios.get.mockResolvedValue(data)
    wrapper = shallowMount(Home, {store, localVue})
    await wrapper.vm.$nextTick()
    wrapper.vm.getIntervalData()
    expect(wrapper.vm.interval).toEqual(15)
  })
  
})