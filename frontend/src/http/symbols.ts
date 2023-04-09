import { InjectionKey, inject } from 'vue'
import { AxiosInstance } from 'axios'

export const AxiosKey: InjectionKey<AxiosInstance> = Symbol('http')

export function injectStrict<T>(key: InjectionKey<T>, fallback?: T) {
    const resolved = inject(key, fallback)
    if (!resolved) {
      throw new Error(`Could not resolve ${key.description}`)
    }
    return resolved
}