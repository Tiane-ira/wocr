export namespace model {
	
	export class ExportField {
	    id: number;
	    name: string;
	    fieldName: string;
	    export: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ExportField(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.fieldName = source["fieldName"];
	        this.export = source["export"];
	    }
	}
	export class OcrParam {
	    ocrPath: string;
	    savePath: string;
	    recursive: boolean;
	    id: string;
	    name: string;
	    type: string;
	    ak: string;
	    sk: string;
	    date: string;
	    isDefault: boolean;
	
	    static createFrom(source: any = {}) {
	        return new OcrParam(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ocrPath = source["ocrPath"];
	        this.savePath = source["savePath"];
	        this.recursive = source["recursive"];
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.ak = source["ak"];
	        this.sk = source["sk"];
	        this.date = source["date"];
	        this.isDefault = source["isDefault"];
	    }
	}
	export class SkConfig {
	    id: string;
	    name: string;
	    type: string;
	    ak: string;
	    sk: string;
	    date: string;
	    isDefault: boolean;
	
	    static createFrom(source: any = {}) {
	        return new SkConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.ak = source["ak"];
	        this.sk = source["sk"];
	        this.date = source["date"];
	        this.isDefault = source["isDefault"];
	    }
	}

}

