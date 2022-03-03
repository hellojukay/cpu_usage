export interface go {
  "main": {
    "CPU": {
		Usage():Promise<string>
    },
  }

}

declare global {
	interface Window {
		go: go;
	}
}
