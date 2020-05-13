import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import InfoBoxPlants from '../src/components/home/InfoBoxPlants.vue'
import vuetify from "vuetify"


describe('InfoxBoxPlants',  () => {
  let wrapper;
  beforeEach(() => {
    const localVue =  createLocalVue()
    localVue.use(vuetify)
    wrapper = shallowMount(InfoBoxPlants, {localVue})    
  })

  it('has a created hook',()  => {
    expect(wrapper.isVueInstance()).toBe(true)  
  })

  it('renders the heading',async()=> {
    wrapper = shallowMount(InfoBoxPlants, {propsData:{
      heading: "TestingD58W"
    }})    
    await wrapper.vm.$nextTick()
    expect(wrapper.html()).toContain('TestingD58W')
  })

  it('renders right amount of plants',async()  => {
    let p1 = {
      plant_type: 'Basil',
      available_plants: 5
    }
    let p2 = {
      plant_type: 'Lettuce',
      available_plants: 0
    }
    wrapper = shallowMount(InfoBoxPlants, {propsData:{
      plants: [p1,p2]
    }})   
    await wrapper.vm.$nextTick()
    expect(wrapper.text()).toContain('Basil 5')
    expect(wrapper.text()).toContain('Lettuce 0')
  })
  
  
})