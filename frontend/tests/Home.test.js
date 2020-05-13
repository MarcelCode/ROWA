import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import vuetify from "vuetify"
import Vuex, { mapState } from 'vuex'
import Module from '../src/views/Home.vue'
import axios from 'axios';

jest.mock('axios');

describe('Home',  () => {
  let wrapper;
  beforeEach(() => {
    let localVue =  createLocalVue()
    localVue.use(vuetify)
    localVue.use(Vuex)
    axios.get.mockResolvedValue({})
    wrapper = shallowMount(Module, {localVue})
  })

  it('has a created hook', async()  => {
    
   // wrapper = shallowMount(Module)
  // expect(wrapper.isVueInstance()).toBe(true)  
  })
  
})