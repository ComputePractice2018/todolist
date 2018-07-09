import { expect } from 'chai'
import { shallowMount } from '@vue/test-utils'
import Dolist from '@/components/Dolist.vue'

describe('Dolist.vue', () => {
  it('renders props.title when passed', () => {
    const title = 'Название'
    const wrapper = shallowMount(Dolist, {
      propsData: { title }
    })
    expect(wrapper.text()).to.include(title)
  })
  it('renders titles', () => {
    const wrapper = shallowMount(Dolist, {})
    expect(wrapper.text()).to.include('ID')
    expect(wrapper.text()).to.include('Название')
    expect(wrapper.text()).to.include('Выполнена')
  })
})
