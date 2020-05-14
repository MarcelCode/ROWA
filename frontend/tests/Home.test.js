import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import vuetify from "vuetify"
import Vuex, { mapState } from 'vuex'
import HomeTopRow from "../src/components/main/HomeTopRow"
import FarmInfo from "../src/components/home/FarmInfo";
import CatTree from "../src/components/home/Farm/CatTree";
import StatGraphic from "../src/components/home/Stats/StatGraphic";
import Home from '../src/views/Home.vue'
import axios from 'axios';

jest.mock('axios');
describe('Home',  () => {
  let wrapper;
  let store;
  let localVue
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
    //expect(wrapper.contains(StatGraphic)).toBe(true)
  })

  it('returns the right sensor data', async()=>{
    let data = {
      data: [{
      datetime: null,
      temperature: 22,
      light_intensity: 150
    }]
    }
    axios.get.mockResolvedValue(data)
    wrapper = shallowMount(Home, {store, localVue})
    await wrapper.vm.$nextTick()
    console.log(data.data[0])
    expect(wrapper.vm.sensor_data.temperature).toEqual(22)
  })
  
})