// Import `shallowMount` from Vue Test Utils and the component being tested
import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import FarmInfo from '../components/home/FarmInfo.vue'
import vuetify from "vuetify"
import axios from 'axios';


jest.mock('axios');
axios.get.mockResolvedValue({})
// Mount the component
//const wrapper = shallowMount(FarmInfo)
// Here are some Jest tests, though you can
// use any test runner/assertion library combo you prefer
describe('FarmInfo',  () => {
  let wrapper;
  beforeEach(() => {
    const localVue =  createLocalVue()
    localVue.use(vuetify)
    wrapper = shallowMount(FarmInfo, {localVue})
    
  })
  
  // Inspect the raw component options
  it('has a created hook',() => {
    expect(shallowMount(FarmInfo).isVueInstance()).toBe(true)
    
  })
})