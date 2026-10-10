export async function loadInit () {
  try {
    const response = await window.client.init({})
    window.initResponse = response
    window.dispatchEvent(new CustomEvent('init-reloaded'))
    return response
  } catch (error) {
    console.error('Failed to initialize:', error)
    window.initResponse = null
    throw error
  }
}
