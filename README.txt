WHAT

        A Prometheus exporter to retrieve data from `/proc/pressure/*`.


REQUIREMENTS

        linux v4.20+


INSTALL

        go get -u -v github.com/cirocosta/psi_exporter


EXAMPLE

        # HELP psi_cpu_avg10 Values for avg10 under /proc/pressure/cpu
        # HELP psi_cpu_avg300 Values for avg300 under /proc/pressure/cpu
        # HELP psi_cpu_avg60 Values for avg60 under /proc/pressure/cpu
        # HELP psi_cpu_total Values for total under /proc/pressure/cpu
        # HELP psi_io_avg10 Values for avg10 under /proc/pressure/io
        # HELP psi_io_avg300 Values for avg300 under /proc/pressure/io
        # HELP psi_io_avg60 Values for avg60 under /proc/pressure/io
        # HELP psi_io_total Values for total under /proc/pressure/io
        # HELP psi_memory_avg10 Values for avg10 under /proc/pressure/memory
        # HELP psi_memory_avg300 Values for avg300 under /proc/pressure/memory
        # HELP psi_memory_avg60 Values for avg60 under /proc/pressure/memory
        # HELP psi_memory_total Values for total under /proc/pressure/memory
        psi_cpu_avg10{type="some"} 0
        psi_cpu_avg300{type="some"} 0.27
        psi_cpu_avg60{type="some"} 0.05
        psi_cpu_total{type="some"} 6.492446e+06
        psi_io_avg10{type="full"} 0
        psi_io_avg10{type="some"} 0
        psi_io_avg300{type="full"} 0.02
        psi_io_avg300{type="some"} 0.03
        psi_io_avg60{type="full"} 0
        psi_io_avg60{type="some"} 0
        psi_io_total{type="full"} 1.124469e+06
        psi_io_total{type="some"} 2.319885e+06
        psi_memory_avg10{type="full"} 0
        psi_memory_avg10{type="some"} 0
        psi_memory_avg300{type="full"} 0
        psi_memory_avg300{type="some"} 0
        psi_memory_avg60{type="full"} 0
        psi_memory_avg60{type="some"} 0
        psi_memory_total{type="full"} 0
        psi_memory_total{type="some"} 0
