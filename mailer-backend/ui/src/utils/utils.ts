export const getInitialsName = (name?: string) => {
  const words = name?.split(' ')
  return `${words?.[0]?.[0].toUpperCase() ?? 'N'}${words?.[1]?.[0].toUpperCase() ?? 'N'}`
}
