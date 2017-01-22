FROM scratch
ADD minimalweb /
EXPOSE 2888
CMD ["/minimalweb"]