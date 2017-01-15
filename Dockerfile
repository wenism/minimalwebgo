FROM scratch
ADD minimalweb /
ADD hello.template.html /
EXPOSE 9999
CMD ["/minimalweb"]