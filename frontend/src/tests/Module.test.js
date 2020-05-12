import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import InfoBoxPlants from '../components/home/InfoBoxPlants.vue'
import vuetify from "vuetify"


describe('Module',  () => {
  let wrapper;
  beforeEach(() => {
    const localVue =  createLocalVue()
    localVue.use(vuetify)
    wrapper = shallowMount(Module, {localVue})    
  })

  it('has a created hook',()  => {
    expect(wrapper.isVueInstance()).toBe(true)  
  })
  it('', ()=>{
      
  })
  
})