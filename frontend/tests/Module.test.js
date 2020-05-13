import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import vuetify from "vuetify"
import Vuex from 'vuex'
import Module from '../src/components/home/Farm/Module.vue'


describe('Module',  () => {
  let wrapper;
  let store;
  let actions;
  let state;
  beforeEach(() => {
    let localVue =  createLocalVue()
    localVue.use(vuetify)
    localVue.use(Vuex)
    actions = {
      change_dash_state: jest.fn(),
      change_bash_state: jest.fn()
    }
    state = {farm_info:{}}
    store = new Vuex.Store({
      state,
      actions,
      getters: Module.getters
    })
    wrapper = shallowMount(Module, {localVue, 
      computed: {
        module_plants: () => {
          return 0
        },
        positions: () => {
          return 0
        }
      }
    })    
  })

  it('has a created hook',()  => {
   expect(wrapper.isVueInstance()).toBe(true)  
  })
  
  it('calculates the right image size',() => {
    expect(wrapper.vm.calculate_img_size(0)).toEqual(10)
    expect(wrapper.vm.calculate_img_size(40)).toEqual(50)
    expect(wrapper.vm.calculate_img_size(21)).toEqual(25)
  })

  
})