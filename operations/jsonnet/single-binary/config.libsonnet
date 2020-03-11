{
    _images+:: {
        tempo: 'joeelliott/canary-frigg:b14df74c',
        tempo_query: 'joeelliott/canary-frigg-query:b14df74c',
    },

    _config+:: {
        port: 3100,
        pvc_size: error 'Must specify a pvc size',
        pvc_storage_class: error 'Must specify a pvc storage class',
        receivers: error 'Must specify receivers',
        ballast_size_mbs: '1024',
        jaeger_ui: {
            base_path: '/',
        }
    },

    tempo_container+::
        $.util.resourcesRequests('3', '3Gi') +
        $.util.resourcesLimits('5', '5Gi')
}